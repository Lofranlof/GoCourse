package tarstream_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"

	"gitlab.com/manytask/itmo-go/public/distbuild/pkg/tarstream"
)

func TestTarStream(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "tarstream")
	require.NoError(t, err)

	t.Logf("running inside %s", tmpDir)

	from := filepath.Join(tmpDir, "from")
	to := filepath.Join(tmpDir, "to")

	require.NoError(t, os.Mkdir(from, 0o777))
	require.NoError(t, os.Mkdir(to, 0o777))

	var buf bytes.Buffer

	require.NoError(t, os.Mkdir(filepath.Join(from, "a"), 0o777))
	require.NoError(t, os.MkdirAll(filepath.Join(from, "b", "c", "d"), 0o777))
	require.NoError(t, ioutil.WriteFile(filepath.Join(from, "a", "x.bin"), []byte("xxx"), 0o777))
	require.NoError(t, ioutil.WriteFile(filepath.Join(from, "b", "c", "y.txt"), []byte("yyy"), 0o666))

	require.NoError(t, tarstream.Send(from, &buf))

	require.NoError(t, tarstream.Receive(to, &buf))

	checkDir := func(path string) {
		st, err := os.Stat(path)
		require.NoError(t, err)
		require.True(t, st.IsDir())
	}

	checkDir(filepath.Join(to, "a"))
	checkDir(filepath.Join(to, "b", "c", "d"))

	checkFile := func(path string, content []byte, mode os.FileMode) {
		t.Helper()

		st, err := os.Stat(path)
		require.NoError(t, err)

		require.Equal(t, mode.String(), st.Mode().String())

		b, err := ioutil.ReadFile(path)
		require.NoError(t, err)
		require.Equal(t, content, b)
	}

	checkFile(filepath.Join(to, "a", "x.bin"), []byte("xxx"), 0o755)
	checkFile(filepath.Join(to, "b", "c", "y.txt"), []byte("yyy"), 0o644)
}

func init() {
	unix.Umask(0o022)
}
