#include "utils.hpp"

std::vector<std::string> SplitByString(std::string s, std::string sep)
{
    StringSlice *slice = Split(s.c_str(), sep.c_str());
    std::vector<std::string> ret;
    for (int i = 0; i < slice->len; i++)
        ret.push_back(slice->data[i]);
    return ret;
}

char *_RunCommand(char *cmd)
{
    return RunCommandWithError(cmd)->out;
}

GSysRes *_RunCommandWithError(char *cmd)
{
    StringSlice *res = Split(cmd, " ");
    char *name = nullptr;
    char **_args_ = nullptr;
    if (res->len > 0)
        name = res->data[0];
    if (res->len > 1)
    {
        _args_ = ++res->data;
        for (int i = 0; i < res->len - 1; i++)
            _args_[i] = ReplaceAll(_args_[i], "&&&", " ");
    }
    GoSlice args = {_args_, res->len - 1, res->len - 1};
    return GSystem(name, args);
}

void RunCmd(char *cmd)
{
    StringSlice *res = Split(cmd, " ");
    char *name = nullptr;
    char **_args_ = nullptr;
    if (res->len > 0)
        name = res->data[0];
    if (res->len > 1)
    {
        _args_ = ++res->data;
        for (int i = 0; i < res->len - 1; i++)
            _args_[i] = ReplaceAll(_args_[i], "&&&", " ");
    }
    GoSlice args = {_args_, res->len - 1, res->len - 1};
    GSystemWithStdout(name, args, 1, 1, "");
}

void Mount(std::string k, std::string v)
{
#ifdef _WIN32
    std::string cmd = FmtCommand("cmd /C setx", (char *)k.c_str(), (char *)v.c_str());
    std::cout << RunCommand(cmd.c_str());
#elif __linux__
    ;
#elif __APPLE__
    ;
#endif
}

std::map<std::string, std::string> InitMapper(std::string workdir)
{
    // 去向daemon发请求拿到路径，读取路径
    std::map<std::string, std::string> mapper;
    return mapper;
}

std::string mapping(std::string args)
{
    if (store->global)
        args = mappingHelper(args, store->gmapper);
    if (store->local)
        args = mappingHelper(args, store->lmapper);
    return args;
}

std::string mappingHelper(std::string args, std::map<std::string, std::string> mapper)
{
    auto iter = mapper.begin();
    while (iter != mapper.end())
    {
        args = ReplaceAll(args.c_str(), iter->first.c_str(), iter->second.c_str());
        iter++;
    }
    return args;
}

std::string _FmtCommand(char *arg, ...)
{
    va_list arg_ptr;
    std::string args = arg;
    char *tmpl;
    va_start(arg_ptr, arg);
    while ((tmpl = va_arg(arg_ptr, char *)) != NULL)
    {
        args += " ";
        args += tmpl;
        arg += strlen(tmpl);
    }
    va_end(arg_ptr);
    return args;
}

bool Conn::Close()
{
    return CloseConn(c) == TRUE ? true : false;
}

std::string Conn::Read()
{
    return ReadConn(c);
}

bool Conn::Write(std::string data)
{
    return WriteConn(c, (char *)data.c_str()) == TRUE ? true : false;
}

void Store::initStore()
{
    global = local = true;
    stop = false;
#ifdef _WIN32
    sys_env = InitWIN32Mapper(QUERY_SYS_ENVIRONMENT, SYS_ENVIRONMENT);
    user_env = InitWIN32Mapper(QUERY_USER_ENVIRONMENT, USER_ENVIRONMENT);
    sys_path = SplitByString(sys_env["Path"].c_str(), ";");
    user_path = SplitByString(user_env["Path"].c_str(), ";");
    sys_env.erase("Path");
    user_env.erase("Path");
#elif __linux__

#elif __APPLE__

#endif
}

std::map<std::string, std::string> InitWIN32Mapper(std::string cmd, std::string cmp)
{
    std::string res = RunCommand(cmd.c_str());
    std::vector<std::string> ret = SplitByString(res.c_str(), "\n");
    std::map<std::string, std::string> env;
    for (int i = 0; i < ret.size(); i++)
    {
        std::string str = ret[i];
        if (str.size() > 1 && FindStr(str.c_str(), cmp.c_str()) < 0)
        {
            std::string tmpl = ReplaceWithRegexp("REG_\\w+", ret[i].c_str(), " ");
            std::string key = "", value = "";
            bool isKey = true;
            for (char ch : tmpl)
            {
                if (ch != ' ' && isKey)
                    key += ch;
                else if (ch != ' ')
                    value += ch;
                if (ch == ' ' && key.size() > 0)
                    isKey = false;
            }
            env[key] = value;
        }
    }
    return env;
}

void Task::Add(std::string cmd)
{
    task_queue.push(cmd);
}

bool Task::Exec()
{
    std::string cmd = task_queue.front();
    GSysRes *res = RunCommandWithError(cmd.c_str());
    if (res->err == TRUE)
    {
        if (store->stop)
        {
            // 要暂停需要看看需不需要回档，同时不pop

            return false;
        }
        glog(WARN, strcat((char *)"Skip a err: ", res->eout));
        task_queue.pop();
        return false;
    }
    return true;
    // 执行有两种模式一种错误停止，一种一直执行下去 ,执行不下去可以调用回档的处理
}

void glog(LOG_TYPE type, std::string data)
{
    std::string v = "";
    switch (type)
    {
    case INFO:
        v += "[INFO] ";
        v += data.c_str();
        _glog("INFO", v.c_str());
        break;
    case WARN:
        v += "[WARN] ";
        v += data.c_str();
        _glog("INFO", v.c_str());
    case FATAL:
        v += "[FATAL] ";
        v += data.c_str();
        _glog("FATAL", data.c_str());
        break;
    case PANIC:
        v += "[PANIC] ";
        v += data.c_str();
        _glog("PANIC", data.c_str());
        break;
    default:
        _glog("INFO", "[WARN] Unknown log type");
        break;
    }
}