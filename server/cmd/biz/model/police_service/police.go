// Code generated by thriftgo (0.3.17). DO NOT EDIT.

package police_service

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type Location struct {
	Latitude  float64 `thrift:"latitude,1" form:"latitude" json:"latitude" query:"latitude"`
	Longitude float64 `thrift:"longitude,2" form:"longitude" json:"longitude" query:"longitude"`
}

func NewLocation() *Location {
	return &Location{}
}

func (p *Location) InitDefault() {
}

func (p *Location) GetLatitude() (v float64) {
	return p.Latitude
}

func (p *Location) GetLongitude() (v float64) {
	return p.Longitude
}

var fieldIDToName_Location = map[int16]string{
	1: "latitude",
	2: "longitude",
}

func (p *Location) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Location[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Location) ReadField1(iprot thrift.TProtocol) error {

	var _field float64
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Latitude = _field
	return nil
}
func (p *Location) ReadField2(iprot thrift.TProtocol) error {

	var _field float64
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Longitude = _field
	return nil
}

func (p *Location) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Location"); err != nil {
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

func (p *Location) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *Location) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *Location) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Location(%+v)", *p)

}

type PoliceStation struct {
	Name     string    `thrift:"name,1" form:"name" json:"name" query:"name"`
	Location *Location `thrift:"location,2" form:"location" json:"location" query:"location"`
	// 距离单位：米
	Distance float64 `thrift:"distance,3" form:"distance" json:"distance" query:"distance"`
}

func NewPoliceStation() *PoliceStation {
	return &PoliceStation{}
}

func (p *PoliceStation) InitDefault() {
}

func (p *PoliceStation) GetName() (v string) {
	return p.Name
}

var PoliceStation_Location_DEFAULT *Location

func (p *PoliceStation) GetLocation() (v *Location) {
	if !p.IsSetLocation() {
		return PoliceStation_Location_DEFAULT
	}
	return p.Location
}

func (p *PoliceStation) GetDistance() (v float64) {
	return p.Distance
}

var fieldIDToName_PoliceStation = map[int16]string{
	1: "name",
	2: "location",
	3: "distance",
}

func (p *PoliceStation) IsSetLocation() bool {
	return p.Location != nil
}

func (p *PoliceStation) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PoliceStation[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PoliceStation) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Name = _field
	return nil
}
func (p *PoliceStation) ReadField2(iprot thrift.TProtocol) error {
	_field := NewLocation()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Location = _field
	return nil
}
func (p *PoliceStation) ReadField3(iprot thrift.TProtocol) error {

	var _field float64
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Distance = _field
	return nil
}

func (p *PoliceStation) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("PoliceStation"); err != nil {
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

func (p *PoliceStation) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Name); err != nil {
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

func (p *PoliceStation) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("location", thrift.STRUCT, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Location.Write(oprot); err != nil {
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

func (p *PoliceStation) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("distance", thrift.DOUBLE, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteDouble(p.Distance); err != nil {
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

func (p *PoliceStation) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PoliceStation(%+v)", *p)

}

type NearestPoliceStationsResponse struct {
	Status string           `thrift:"status,1" form:"status" json:"status" query:"status"`
	Data   []*PoliceStation `thrift:"data,2" form:"data" json:"data" query:"data"`
}

func NewNearestPoliceStationsResponse() *NearestPoliceStationsResponse {
	return &NearestPoliceStationsResponse{}
}

func (p *NearestPoliceStationsResponse) InitDefault() {
}

func (p *NearestPoliceStationsResponse) GetStatus() (v string) {
	return p.Status
}

func (p *NearestPoliceStationsResponse) GetData() (v []*PoliceStation) {
	return p.Data
}

var fieldIDToName_NearestPoliceStationsResponse = map[int16]string{
	1: "status",
	2: "data",
}

func (p *NearestPoliceStationsResponse) Read(iprot thrift.TProtocol) (err error) {

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
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.LIST {
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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_NearestPoliceStationsResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *NearestPoliceStationsResponse) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Status = _field
	return nil
}
func (p *NearestPoliceStationsResponse) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	_field := make([]*PoliceStation, 0, size)
	values := make([]PoliceStation, size)
	for i := 0; i < size; i++ {
		_elem := &values[i]
		_elem.InitDefault()

		if err := _elem.Read(iprot); err != nil {
			return err
		}

		_field = append(_field, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	p.Data = _field
	return nil
}

func (p *NearestPoliceStationsResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("NearestPoliceStationsResponse"); err != nil {
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

func (p *NearestPoliceStationsResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("status", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Status); err != nil {
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

func (p *NearestPoliceStationsResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("data", thrift.LIST, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Data)); err != nil {
		return err
	}
	for _, v := range p.Data {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
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

func (p *NearestPoliceStationsResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("NearestPoliceStationsResponse(%+v)", *p)

}

type PoliceService interface {
	GetNearestPoliceStations(ctx context.Context, userLocation *Location) (r *NearestPoliceStationsResponse, err error)
}

type PoliceServiceClient struct {
	c thrift.TClient
}

func NewPoliceServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PoliceServiceClient {
	return &PoliceServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPoliceServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PoliceServiceClient {
	return &PoliceServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPoliceServiceClient(c thrift.TClient) *PoliceServiceClient {
	return &PoliceServiceClient{
		c: c,
	}
}

func (p *PoliceServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *PoliceServiceClient) GetNearestPoliceStations(ctx context.Context, userLocation *Location) (r *NearestPoliceStationsResponse, err error) {
	var _args PoliceServiceGetNearestPoliceStationsArgs
	_args.UserLocation = userLocation
	var _result PoliceServiceGetNearestPoliceStationsResult
	if err = p.Client_().Call(ctx, "getNearestPoliceStations", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type PoliceServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      PoliceService
}

func (p *PoliceServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *PoliceServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *PoliceServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewPoliceServiceProcessor(handler PoliceService) *PoliceServiceProcessor {
	self := &PoliceServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("getNearestPoliceStations", &policeServiceProcessorGetNearestPoliceStations{handler: handler})
	return self
}
func (p *PoliceServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
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

type policeServiceProcessorGetNearestPoliceStations struct {
	handler PoliceService
}

func (p *policeServiceProcessorGetNearestPoliceStations) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PoliceServiceGetNearestPoliceStationsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getNearestPoliceStations", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := PoliceServiceGetNearestPoliceStationsResult{}
	var retval *NearestPoliceStationsResponse
	if retval, err2 = p.handler.GetNearestPoliceStations(ctx, args.UserLocation); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getNearestPoliceStations: "+err2.Error())
		oprot.WriteMessageBegin("getNearestPoliceStations", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getNearestPoliceStations", thrift.REPLY, seqId); err2 != nil {
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

type PoliceServiceGetNearestPoliceStationsArgs struct {
	UserLocation *Location `thrift:"userLocation,1"`
}

func NewPoliceServiceGetNearestPoliceStationsArgs() *PoliceServiceGetNearestPoliceStationsArgs {
	return &PoliceServiceGetNearestPoliceStationsArgs{}
}

func (p *PoliceServiceGetNearestPoliceStationsArgs) InitDefault() {
}

var PoliceServiceGetNearestPoliceStationsArgs_UserLocation_DEFAULT *Location

func (p *PoliceServiceGetNearestPoliceStationsArgs) GetUserLocation() (v *Location) {
	if !p.IsSetUserLocation() {
		return PoliceServiceGetNearestPoliceStationsArgs_UserLocation_DEFAULT
	}
	return p.UserLocation
}

var fieldIDToName_PoliceServiceGetNearestPoliceStationsArgs = map[int16]string{
	1: "userLocation",
}

func (p *PoliceServiceGetNearestPoliceStationsArgs) IsSetUserLocation() bool {
	return p.UserLocation != nil
}

func (p *PoliceServiceGetNearestPoliceStationsArgs) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PoliceServiceGetNearestPoliceStationsArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PoliceServiceGetNearestPoliceStationsArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewLocation()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.UserLocation = _field
	return nil
}

func (p *PoliceServiceGetNearestPoliceStationsArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("getNearestPoliceStations_args"); err != nil {
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

func (p *PoliceServiceGetNearestPoliceStationsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("userLocation", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.UserLocation.Write(oprot); err != nil {
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

func (p *PoliceServiceGetNearestPoliceStationsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PoliceServiceGetNearestPoliceStationsArgs(%+v)", *p)

}

type PoliceServiceGetNearestPoliceStationsResult struct {
	Success *NearestPoliceStationsResponse `thrift:"success,0,optional"`
}

func NewPoliceServiceGetNearestPoliceStationsResult() *PoliceServiceGetNearestPoliceStationsResult {
	return &PoliceServiceGetNearestPoliceStationsResult{}
}

func (p *PoliceServiceGetNearestPoliceStationsResult) InitDefault() {
}

var PoliceServiceGetNearestPoliceStationsResult_Success_DEFAULT *NearestPoliceStationsResponse

func (p *PoliceServiceGetNearestPoliceStationsResult) GetSuccess() (v *NearestPoliceStationsResponse) {
	if !p.IsSetSuccess() {
		return PoliceServiceGetNearestPoliceStationsResult_Success_DEFAULT
	}
	return p.Success
}

var fieldIDToName_PoliceServiceGetNearestPoliceStationsResult = map[int16]string{
	0: "success",
}

func (p *PoliceServiceGetNearestPoliceStationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PoliceServiceGetNearestPoliceStationsResult) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PoliceServiceGetNearestPoliceStationsResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PoliceServiceGetNearestPoliceStationsResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewNearestPoliceStationsResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *PoliceServiceGetNearestPoliceStationsResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("getNearestPoliceStations_result"); err != nil {
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

func (p *PoliceServiceGetNearestPoliceStationsResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *PoliceServiceGetNearestPoliceStationsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PoliceServiceGetNearestPoliceStationsResult(%+v)", *p)

}