// Code generated by thriftgo (0.3.17). DO NOT EDIT.

package urban

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// 请求结构体
type UrbanHeatIslandRequest struct {
	// 中心点纬度
	Latitude float64 `thrift:"latitude,1" form:"latitude" json:"latitude" query:"latitude"`
	// 中心点经度
	Longitude float64 `thrift:"longitude,2" form:"longitude" json:"longitude" query:"longitude"`
	// 半径（米）
	Radius float64 `thrift:"radius,3" form:"radius" json:"radius" query:"radius"`
}

func NewUrbanHeatIslandRequest() *UrbanHeatIslandRequest {
	return &UrbanHeatIslandRequest{}
}

func (p *UrbanHeatIslandRequest) InitDefault() {
}

func (p *UrbanHeatIslandRequest) GetLatitude() (v float64) {
	return p.Latitude
}

func (p *UrbanHeatIslandRequest) GetLongitude() (v float64) {
	return p.Longitude
}

func (p *UrbanHeatIslandRequest) GetRadius() (v float64) {
	return p.Radius
}

var fieldIDToName_UrbanHeatIslandRequest = map[int16]string{
	1: "latitude",
	2: "longitude",
	3: "radius",
}

func (p *UrbanHeatIslandRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.DOUBLE {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.DOUBLE {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.DOUBLE {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UrbanHeatIslandRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UrbanHeatIslandRequest) ReadField1(iprot thrift.TProtocol) error {

	var _field float64
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Latitude = _field
	return nil
}
func (p *UrbanHeatIslandRequest) ReadField2(iprot thrift.TProtocol) error {

	var _field float64
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Longitude = _field
	return nil
}
func (p *UrbanHeatIslandRequest) ReadField3(iprot thrift.TProtocol) error {

	var _field float64
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Radius = _field
	return nil
}

func (p *UrbanHeatIslandRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UrbanHeatIslandRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UrbanHeatIslandRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("latitude", thrift.DOUBLE, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteDouble(p.Latitude); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UrbanHeatIslandRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("longitude", thrift.DOUBLE, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteDouble(p.Longitude); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UrbanHeatIslandRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("radius", thrift.DOUBLE, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteDouble(p.Radius); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *UrbanHeatIslandRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UrbanHeatIslandRequest(%+v)", *p)

}

// 响应结构体
type UrbanHeatIslandResponse struct {
	// 城市热岛效应指数
	HeatIslandIndex float64 `thrift:"heatIslandIndex,1" form:"heatIslandIndex" json:"heatIslandIndex" query:"heatIslandIndex"`
	// 土地利用类型 -> 面积
	LandUseAreas map[string]float64 `thrift:"landUseAreas,2" form:"landUseAreas" json:"landUseAreas" query:"landUseAreas"`
}

func NewUrbanHeatIslandResponse() *UrbanHeatIslandResponse {
	return &UrbanHeatIslandResponse{}
}

func (p *UrbanHeatIslandResponse) InitDefault() {
}

func (p *UrbanHeatIslandResponse) GetHeatIslandIndex() (v float64) {
	return p.HeatIslandIndex
}

func (p *UrbanHeatIslandResponse) GetLandUseAreas() (v map[string]float64) {
	return p.LandUseAreas
}

var fieldIDToName_UrbanHeatIslandResponse = map[int16]string{
	1: "heatIslandIndex",
	2: "landUseAreas",
}

func (p *UrbanHeatIslandResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.DOUBLE {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.MAP {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UrbanHeatIslandResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UrbanHeatIslandResponse) ReadField1(iprot thrift.TProtocol) error {

	var _field float64
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		_field = v
	}
	p.HeatIslandIndex = _field
	return nil
}
func (p *UrbanHeatIslandResponse) ReadField2(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return err
	}
	_field := make(map[string]float64, size)
	for i := 0; i < size; i++ {
		var _key string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_key = v
		}

		var _val float64
		if v, err := iprot.ReadDouble(); err != nil {
			return err
		} else {
			_val = v
		}

		_field[_key] = _val
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return err
	}
	p.LandUseAreas = _field
	return nil
}

func (p *UrbanHeatIslandResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UrbanHeatIslandResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UrbanHeatIslandResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("heatIslandIndex", thrift.DOUBLE, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteDouble(p.HeatIslandIndex); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UrbanHeatIslandResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("landUseAreas", thrift.MAP, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.DOUBLE, len(p.LandUseAreas)); err != nil {
		return err
	}
	for k, v := range p.LandUseAreas {
		if err := oprot.WriteString(k); err != nil {
			return err
		}
		if err := oprot.WriteDouble(v); err != nil {
			return err
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UrbanHeatIslandResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UrbanHeatIslandResponse(%+v)", *p)

}

// 服务接口
type UrbanHeatIslandService interface {
	AnalyzeUrbanHeatIsland(ctx context.Context, request *UrbanHeatIslandRequest) (r *UrbanHeatIslandResponse, err error)
}

type UrbanHeatIslandServiceClient struct {
	c thrift.TClient
}

func NewUrbanHeatIslandServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *UrbanHeatIslandServiceClient {
	return &UrbanHeatIslandServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewUrbanHeatIslandServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *UrbanHeatIslandServiceClient {
	return &UrbanHeatIslandServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewUrbanHeatIslandServiceClient(c thrift.TClient) *UrbanHeatIslandServiceClient {
	return &UrbanHeatIslandServiceClient{
		c: c,
	}
}

func (p *UrbanHeatIslandServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *UrbanHeatIslandServiceClient) AnalyzeUrbanHeatIsland(ctx context.Context, request *UrbanHeatIslandRequest) (r *UrbanHeatIslandResponse, err error) {
	var _args UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs
	_args.Request = request
	var _result UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult
	if err = p.Client_().Call(ctx, "AnalyzeUrbanHeatIsland", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type UrbanHeatIslandServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      UrbanHeatIslandService
}

func (p *UrbanHeatIslandServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *UrbanHeatIslandServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *UrbanHeatIslandServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewUrbanHeatIslandServiceProcessor(handler UrbanHeatIslandService) *UrbanHeatIslandServiceProcessor {
	self := &UrbanHeatIslandServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("AnalyzeUrbanHeatIsland", &urbanHeatIslandServiceProcessorAnalyzeUrbanHeatIsland{handler: handler})
	return self
}
func (p *UrbanHeatIslandServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type urbanHeatIslandServiceProcessorAnalyzeUrbanHeatIsland struct {
	handler UrbanHeatIslandService
}

func (p *urbanHeatIslandServiceProcessorAnalyzeUrbanHeatIsland) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("AnalyzeUrbanHeatIsland", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult{}
	var retval *UrbanHeatIslandResponse
	if retval, err2 = p.handler.AnalyzeUrbanHeatIsland(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing AnalyzeUrbanHeatIsland: "+err2.Error())
		oprot.WriteMessageBegin("AnalyzeUrbanHeatIsland", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("AnalyzeUrbanHeatIsland", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs struct {
	Request *UrbanHeatIslandRequest `thrift:"request,1"`
}

func NewUrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs() *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs {
	return &UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs{}
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs) InitDefault() {
}

var UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs_Request_DEFAULT *UrbanHeatIslandRequest

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs) GetRequest() (v *UrbanHeatIslandRequest) {
	if !p.IsSetRequest() {
		return UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs_Request_DEFAULT
	}
	return p.Request
}

var fieldIDToName_UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs = map[int16]string{
	1: "request",
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewUrbanHeatIslandRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Request = _field
	return nil
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("AnalyzeUrbanHeatIsland_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UrbanHeatIslandServiceAnalyzeUrbanHeatIslandArgs(%+v)", *p)

}

type UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult struct {
	Success *UrbanHeatIslandResponse `thrift:"success,0,optional"`
}

func NewUrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult() *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult {
	return &UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult{}
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult) InitDefault() {
}

var UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult_Success_DEFAULT *UrbanHeatIslandResponse

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult) GetSuccess() (v *UrbanHeatIslandResponse) {
	if !p.IsSetSuccess() {
		return UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult = map[int16]string{
	0: "success",
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewUrbanHeatIslandResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("AnalyzeUrbanHeatIsland_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UrbanHeatIslandServiceAnalyzeUrbanHeatIslandResult(%+v)", *p)

}
