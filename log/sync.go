package log

import (
	"github.com/veerdone/gopkg/conf"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type getSyncFunc func() io.WriteCloser

var syncFuncMap = make(map[string]getSyncFunc)

var wc []io.WriteCloser

func initSyncFuncMap() {
	syncFuncMap["file"] = fileOut
	syncFuncMap["stdout"] = stdout
}

func stdout() io.WriteCloser {
	return os.Stdout
}

func getSyncs(conf conf.Log) []zapcore.WriteSyncer {
	output := conf.Output
	if len(output) == 0 {
		return []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}
	}
	ws := make([]zapcore.WriteSyncer, 0, len(output))
	wc = make([]io.WriteCloser, 0, len(output))
	for i := range output {
		syncFunc, ok := syncFuncMap[output[i]]
		if ok {
			closer := syncFunc()
			wc = append(wc, closer)
			ws = append(ws, zapcore.AddSync(closer))
		}
	}

	return ws
}

func Close() {
	for i := range wc {
		wc[i].Close()
	}
}
