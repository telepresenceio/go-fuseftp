package fs

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/winfsp/cgofuse/fuse"

	"github.com/telepresenceio/clog"
	"github.com/telepresenceio/clog/log"
)

// FuseHost wraps a fuse.FileSystemHost and adds Start/Stop semantics
type FuseHost struct {
	host       *fuse.FileSystemHost
	mountPoint string
	cancel     context.CancelFunc
	wg         sync.WaitGroup
}

// NewHost creates a FuseHost instance that will mount the given filesystem
// on the given mountPoint.
func NewHost(fsh fuse.FileSystemInterface, mountPoint string) *FuseHost {
	host := fuse.NewFileSystemHost(fsh)
	host.SetCapReaddirPlus(true)
	return &FuseHost{host: host, mountPoint: mountPoint}
}

// Start will mount the filesystem on the mountPoint passed to NewHost.
func (fh *FuseHost) Start(ctx context.Context, startTimeout time.Duration) error {
	ctx, fh.cancel = context.WithCancel(ctx)

	opts := []string{
		"-o", "default_permissions",
		"-o", "auto_cache",
		"-o", "sync_read",
		"-o", "allow_root",
	}
	if log.Enabled(clog.LevelTrace) {
		opts = append(opts, "-o", "debug")
	}
	if runtime.GOOS == "windows" {
		// WinFsp requires this to create files with the same
		// user as the one that starts the FUSE mount
		opts = append(opts, "-o", "uid=-1", "-o", "gid=-1")
	}
	started := make(chan error, 1)
	startCtx, startCancel := context.WithTimeout(ctx, startTimeout)
	defer startCancel()
	go fh.detectFuseStarted(startCtx, started)

	mCh := make(chan bool, 1)
	go func() {
		log.Debugf("FuseHost mounting %s", fh.mountPoint)
		mountResult := fh.host.Mount(fh.mountPoint, opts)
		mCh <- mountResult
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		_, _ = os.Stat(fh.mountPoint)
	}()

	fh.wg.Add(1)
	go func() {
		defer fh.wg.Done()
		<-ctx.Done()
		slog.Debug("FuseHost unmounting")
		fh.host.Unmount()
		slog.Debug("FuseHost unmount complete")
		select {
		case mountResult := <-mCh:
			if !mountResult {
				log.Errorf("FuseHost mount of %s failed", fh.mountPoint)
			} else {
				log.Debugf("FuseHost mount of %s ended", fh.mountPoint)
			}
		case <-time.After(5 * time.Second):
			// Mount seem to be stuck.
			log.Errorf("fuse mount of %s timed out", fh.mountPoint)
		}
	}()
	return <-started
}

// Stop will unmount the file system and terminate the FTP client, wait for all clean-up to
// complete, and then return
func (fh *FuseHost) Stop() {
	// cancel will cause the host to unmount, which in turn will result in a call to
	// Destroy() which will terminate the FTP client.
	fh.cancel()
	fh.wg.Wait()
}
