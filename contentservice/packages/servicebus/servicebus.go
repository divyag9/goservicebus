package servicebus

// Request to send to servicebus call
type Request struct {
	Jsonrpc        string  `json:"jsonrpc,omitempty"`
	Method         string  `json:"method,omitempty"`
	Params         *Params `json:"params,omitempty"`
	ID             int32   `json:"id,omitempty"`
	Asyncmessageid int32   `json:"asyncmessageid,omitempty"`
	Traceid        int32   `json:"traceid,omitempty"`
}

// Params are part of the request sent to servicebus call
type Params struct {
	Contractorid int64  `json:"contractorid,omitempty"`
	Ordernumber  int64  `json:"ordernumber,omitempty"`
	Imagetype    int32  `json:"imagetype,omitempty"`
	Filename     string `json:"filename,omitempty"`
	Imagewidth   int32  `json:"imagewidth,omitempty"`
	Imageheight  int32  `json:"imageheight,omitempty"`
	Releasedate  string `json:"releasedate,omitempty"`
	Deptcode     string `json:"deptcode,omitempty"`
	Filecontents []byte `json:"filecontents,omitempty"`
}

// Response from servicebus
type Response struct {
	Jsonrpc string  `json:"jsonrpc,omitempty"`
	ID      int32   `json:"id,omitempty"`
	Result  *Result `json:"result,omitempty"`
	Error   *Error  `json:"error,omitempty"`
}

// Result part of the response
type Result struct {
	Contractorid          int32                 `json:"contractorid,omitempty"`
	Releasedate           string                `json:"releasedate,omitempty"`
	Scandate              string                `json:"scandate,omitempty"`
	Imagetype             int32                 `json:"imagetype,omitempty"`
	Imagewidth            int32                 `json:"imagewidth,omitempty"`
	Imageheight           int32                 `json:"imageheight,omitempty"`
	Deptcode              string                `json:"deptcode,omitempty"`
	Descprefix            string                `json:"descprefix,omitempty"`
	Desctext              string                `json:"desctext,omitempty"`
	Category              string                `json:"category,omitempty"`
	Ordernumber           int64                 `json:"ordernumber,omitempty"`
	Archived              string                `json:"archived,omitempty"`
	Datecreated           string                `json:"datecreated,omitempty"`
	Datemodefied          string                `json:"datemodefied,omitempty"`
	Filesize              int32                 `json:"filesize,omitempty"`
	ID                    int32                 `json:"id,omitempty"`
	Imagefilename         string                `json:"imagefilename,omitempty"`
	Imagerotated          int32                 `json:"imagerotated,omitempty"`
	Thumbnailsize         int32                 `json:"thumbnailsize,omitempty"`
	Webfilename           string                `json:"webfilename,omitempty"`
	Mimetype              string                `json:"mimetype,omitempty"`
	Inspiresponsedata     *InspiPutResponse     `json:"inspiresponsedata,omitempty"`
	Vendorwebresponsedata *VendorWebPutResponse `json:"vendorwebresponsedata,omitempty"`
	GUID                  string                `json:"guid,omitempty"`
}

// VendorWebPutResponse part of result
type VendorWebPutResponse struct {
	Documentid   int64 `json:"documentid,omitempty"`
	Annotationid int64 `json:"annotationid,omitempty"`
}

//InspiPutResponse part of result
type InspiPutResponse struct {
	Photodetailid int64 `json:"photodetailid,omitempty"`
}

// Error part of the response
type Error struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}
