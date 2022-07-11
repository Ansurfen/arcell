#ifndef __GO_UTILS_HPP__
#define __GO_UTILS_HPP__

#include "go/go_utils.h"
#include <string>
#include <vector>
#include <map>
#include <queue>
#include <cstdarg>
#include <stdio.h>
#include <iostream>
#include <string.h>

#define QUERY_USER_ENVIRONMENT (char *)"cmd /C REG QUERY HKEY_CURRENT_USER\\Environment"
#define QUERY_SYS_ENVIRONMENT (char *)"cmd /C REG QUERY HKLM\\SYSTEM\\CurrentControlSet\\Control\\Session&&&Manager\\Environment"
#define USER_ENVIRONMENT (char *)"HKEY_CURRENT_USER\\Environment"
#define SYS_ENVIRONMENT (char *)"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment"

#define _glog(type, data) _log((char *)type, (char *)data);
#define Replace(s, _old, _new, n) _Replace((char *)s, (char *)_old, (char *)_new, n)
#define ReplaceAll(s, _old, _new) Replace(s, _old, _new, -1)
#define Split(s, sep) _Split((char *)s, (char *)sep)
#define ReplaceWithRegexp(expr, src, repl) _ReplaceWithRegexp((char *)expr, (char *)src, (char *)repl)
#define FindStr(s, substr) _FindStr((char *)s, (char *)substr)
#define RunCommand(cmd) _RunCommand((char *)cmd)
#define RunCommandWithError(cmd) _RunCommandWithError((char *)cmd)
#define FmtCommand(arg, ...) _FmtCommand((char *)arg, __VA_ARGS__, NULL)

typedef enum
{
    INFO,
    WARN,
    FATAL,
    PANIC
} LOG_TYPE;

typedef enum
{
    FALSE,
    TRUE
} BOOLEAN;

void RunCmd(char*);
std::string mapping(std::string);
std::string mappingHelper(std::string, std::map<std::string, std::string>);
// 登记
void Reg();
void Unreg();
// 挂载变量
void Mount(std::string, std::string);
void Unmount();
void CopyEnv();
void glog(LOG_TYPE, std::string);
char *_RunCommand(char *);
GSysRes *_RunCommandWithError(char *);
std::vector<std::string> SplitByString(std::string, std::string);
std::map<std::string, std::string> InitMapper(std::string);
std::map<std::string, std::string> InitWIN32Mapper(std::string, std::string);
std::string _FmtCommand(char *arg, ...);

class Store
{
public:
    bool global, local;
    bool stop;
    std::string workdir;
    std::map<std::string, std::string> gmapper, lmapper;
#ifdef _WIN32
    std::map<std::string, std::string> user_env, sys_env;
    std::vector<std::string> user_path, sys_path;
#elif __linux__
    ;
#elif __APPLE__
    ;
#endif
    Store()
    {
        initStore();
    }
    void initStore();
};

class Conn
{
private:
    _Conn *c;

public:
    Conn(std::string network, std::string addr)
    {
        c = DialConn((char *)network.c_str(), (char *)addr.c_str());
    }
    bool Close();
    bool Write(std::string);
    std::string Read();
};

class Task
{
private:
    std::queue<std::string> task_queue;

public:
    void Add(std::string);
    bool Exec();
    bool Run(); //运行整个生命周期
};

Store *store = new Store();

#endif