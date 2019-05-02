package vsz

// API Version: v8_0

type WSGUploadStatisticstoFTPService struct {
    c *Client
}

func NewWSGUploadStatisticstoFTPService (c *Client) *WSGUploadStatisticstoFTPService {
    s := new(WSGUploadStatisticstoFTPService)
    s.c = c
    return s
}

