package models

type Record struct {
	Message        string `parquet:"name=Message, type=BYTE_ARRAY, convertedtype=UTF8"`
	MessageRaw     string `parquet:"name=MessageRaw, type=BYTE_ARRAY, convertedtype=UTF8"`
	StructuredData string `parquet:"name=StructuredData, type=BYTE_ARRAY, convertedtype=UTF8"`
	Tag            string `parquet:"name=Tag, type=BYTE_ARRAY, convertedtype=UTF8"`
	Sender         string `parquet:"name=Sender, type=BYTE_ARRAY, convertedtype=UTF8"`
	Groupings      string `parquet:"name=Groupings, type=BYTE_ARRAY, convertedtype=UTF8"`
	Event          string `parquet:"name=Event, type=BYTE_ARRAY, convertedtype=UTF8"`
	EventId        string `parquet:"name=EventId, type=BYTE_ARRAY, convertedtype=UTF8"`
	NanoTimeStamp  string `parquet:"name=NanoTimeStamp, type=BYTE_ARRAY, convertedtype=UTF8"`
	Namespace      string `parquet:"name=namespace, type=BYTE_ARRAY, convertedtype=UTF8"`
}
