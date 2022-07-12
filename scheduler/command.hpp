#ifndef __AC_COMMAND_HPP_
#define __AC_COMMAND_HPP_
#define _CRT_SECURE_NO_WARNINGS
#include <iostream>
#include <string.h>
#include <string>
#include <string.h>
#include <gflags/gflags.h>  
#include <vector>
#include <fstream>
DEFINE_bool(mount, false, "");
DEFINE_bool(unmount, false, "");
DEFINE_bool(reg, false, "");
DEFINE_bool(unreg, false, "");
DEFINE_bool(map, false, "");
DEFINE_bool(unmap, false, "");
DEFINE_bool(ls, false, "");
DEFINE_bool(cp, false, "");
/// <summary>
/// Ԥ����
/// </summary>
DEFINE_bool(global_mapper, false, "");
DEFINE_bool(local_mapper, false, "");
DEFINE_int32(port, 6379, "what is the port ?");

void ValidatorRegister();
void CheckUniqueFlag();
char** mapper(int,char**);

static bool ValidatePort(const char*, int32_t);

/// <summary>
/// ע�����ص�����,����
/// </summary>
class Command {
private:
	int argc;
	char** argv;

	void mount();
	void unmount();
	void reg();
	void unreg();
	void map();
	void unmap();
	void ls();
	void cp();
public:
	Command(int _argc,char** _argv) {
		argc = _argc;
		argv = mapper(_argc,_argv);
		ValidatorRegister();
	}
	void Excute();
	void ExcuteCmd();
	void ShutDown();
};

void Command::Excute() {
	google::ParseCommandLineFlags(&argc, &argv, true);
	CheckUniqueFlag();
	ExcuteCmd();
}

void Command::ExcuteCmd() {
	if (FLAGS_mount)
		mount();
	if (FLAGS_unmount)
		unmount();
	if (FLAGS_map)
		map();
	if (FLAGS_unmap)
		unmap();
	if (FLAGS_reg)
		reg();
	if (FLAGS_unreg)
		unreg();
	if (FLAGS_cp)
		cp();
	if (FLAGS_ls)
		ls();
}

void Command::ShutDown() {
	google::ShutDownCommandLineFlags();
}

void ValidatorRegister() {
	google::RegisterFlagValidator(&FLAGS_port, &ValidatePort);

}

void CheckUniqueFlag() {
	int BOOLEAN = FLAGS_map + FLAGS_mount + FLAGS_reg + FLAGS_unmap + FLAGS_unmount + FLAGS_unreg;
	if (BOOLEAN > 0 && BOOLEAN <=1) {
		std::cout << BOOLEAN << std::endl;
		return;
	}
	std::cout << "Command isn't only or not found." << std::endl;
}

static bool ValidatePort(const char* flagname, int32_t val)
{
	return val > 0 && val < 32768 ? true : false;
}

void Command::mount() {
	std::cout << argc << std::endl;
	if (argc >= 3) {
		std::cout << argv[1] << ":" << argv[2] << std::endl;
	}
}

void Command::unmount() {
}

void Command::map() {
}

void Command::unmap() {

}

void Command::reg() {

}

void Command::unreg() {

}

void Command::ls() {

}

void Command::cp() {

}

char** mapper(int argc,char** args) {
	
	for (int i = 1; i < argc; i++) {
		args[i];
	}
	return args;
}

#endif // !__COMMAND_HPP_