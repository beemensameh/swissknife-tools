package swissfiles

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	path    string
	FileCmd = &cobra.Command{
		Use:   "file",
		Short: "Use file tool",
	}
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve static files to network",
		Long:  "Open a server and service a specific path to network",
		RunE:  serveStaticAction,
	}
	serveCmd.Flags().StringVarP(&path, "path", "p", "", "The path for the file")
	if err := serveCmd.MarkFlagRequired("path"); err != nil {
		log.Fatal(err)
	}
	FileCmd.AddCommand(serveCmd)
}

func serveStaticAction(cmd *cobra.Command, args []string) error {
	return serveStatic(&ServeStatic{Path: path})
}

func serveStatic(s *ServeStatic) error {
	if err := s.Validated(); err != nil {
		return err
	}
	fs := http.FileServer(http.Dir(s.Path))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("uri=%s header=%s ip=%s", r.RequestURI, r.Header, r.RemoteAddr)
		fs.ServeHTTP(w, r)
	})

	log.Println("Server started at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
	return nil
}
