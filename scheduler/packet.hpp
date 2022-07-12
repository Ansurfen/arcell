#ifndef __AC_PACKET_HPP_
#define __AC_PACKET_HPP_
#include <string>
#include <map>
#include "json/json.h"

std::string serializeHelper(Json::Value);

class Packet {
private:
	std::string _ACTION_;
	std::string _PARAM_;
	std::string _DATA_;
public:
	Packet(std::string action,std::string params,std::string data):
		_ACTION_(action),_PARAM_(params),_DATA_(data) {}
	Packet(std::string action,std::string params,std::map<std::string,std::string> data):
		_ACTION_(action),_PARAM_(params)
	{
		Json::Value tmpl;
		for (std::map<std::string, std::string>::const_iterator iter = data.begin();iter!=data.end();++iter) {
			tmpl[iter->first] = iter->second;
		}
		_DATA_ = serializeHelper(tmpl);
	}
	std::string serialize();
};

std::string Packet::serialize() {
	Json::Value ret;
	ret["_ACTION_"] = _ACTION_;
	ret["_PARAM_"] = _PARAM_;
	ret["_DATA_"] = _DATA_;
	return serializeHelper(ret);
}

std::string serializeHelper(Json::Value json) {
	Json::StreamWriterBuilder writerBuilder;
	writerBuilder.settings_["indentation"] = "";
	std::ostringstream os;
	std::unique_ptr<Json::StreamWriter> jsonWriter(writerBuilder.newStreamWriter());
	jsonWriter->write(json, &os);
	return os.str();
}

#endif // !__AC_PACKET_HPP_
