package server

import (
	"fmt"
	"net/http"
	"strings"
)

func (s *HTTPServer) SnapshotSpecificRequest(resp http.ResponseWriter, req *http.Request) (interface{}, error) {

	fmt.Println("[DEBUG] Processing", req.Method, "request")

	switch req.Method {
	case "PUT", "POST":
		return s.SnapshotCreate(resp, req)
	case "GET":
		return s.SnapshotSpecificGetRequest(resp, req)
	default:
		return nil, CodedError(405, ErrInvalidMethod)
	}

}

// SnapshotSpecificGetRequest deals with HTTP GET request w.r.t a Volume Snapshot
func (s *HTTPServer) SnapshotSpecificGetRequest(resp http.ResponseWriter, req *http.Request) (interface{}, error) {
	// Extract info from path after trimming
	path := strings.TrimPrefix(req.URL.Path, "/latest/snapshot")

	// Is req valid ?
	if path == req.URL.Path {
		fmt.Println("Request coming", path)
		return nil, CodedError(405, ErrInvalidMethod)
	}

	switch {

	case strings.Contains(path, "/revert/"):
		volName := strings.TrimPrefix(path, "/revert/")
		return s.SnapshotRevert(resp, req, volName)
		/*	case strings.Contains(path, "/delete/"):
			snapName := strings.TrimPrefix(path, "/delete/")
			return s.SnapshotDelete(resp, req, volName)
		*/case path == "/list":
		volName := strings.TrimPrefix(path, "/list/")
		return s.SnapshotList(resp, req, volName)
	default:
		return nil, CodedError(405, ErrInvalidMethod)
	}
}

func (s *HTTPServer) SnapshotCreate(resp http.ResponseWriter, req *http.Request) (interface{}, error) {
	volName := strings.TrimPrefix(path, "/latest/snapshot")

	fmt.Println("Volume Name :", volName)
	fmt.Println("[DEBUG] Processing snapshot-create request of volume %s", volName)
	details, err := s.vsmRead(resp, req, volName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Details are :", details)

	return details, nil
}

func (s *HTTPServer) SnapshotRevert(resp http.ResponseWriter, req *http.Request, volName string) (interface{}, error) {

	fmt.Println("Not Implemented")
	return "Not_Implemented", nil
}

func (s *HTTPServer) SnapshotList(resp http.ResponseWriter, req *http.Request, volName string) (interface{}, error) {

	fmt.Println("Not Implemented")
	return "Not_Implemented", nil
}
