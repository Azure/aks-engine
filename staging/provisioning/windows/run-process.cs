using System;
using System.Diagnostics;
using System.IO;
using System.Threading;

namespace RunProcess
{
    public static class exec
    {
        public static void StdOutHandler(object sender, DataReceivedEventArgs e) {
            Console.WriteLine(e.Data);
        }

        public static void StdErrHandler(object sender, DataReceivedEventArgs e) {
            Console.Error.WriteLine(e.Data);
        }
        
        // RunProcess starts a process, sets the priority class of that process to 'High'
        // and redirects all standard output / error for logging.
        public static int RunProcess(string executable, string args = "", ProcessPriorityClass priorityClass = ProcessPriorityClass.Normal) {
            Process p = new Process();
            p.StartInfo.FileName = executable;
            p.StartInfo.WorkingDirectory = Path.GetDirectoryName(executable);
            p.StartInfo.UseShellExecute = false;
            p.StartInfo.CreateNoWindow = true;
            p.StartInfo.RedirectStandardOutput = true;
            p.StartInfo.RedirectStandardError = true;

            if (!String.IsNullOrEmpty(args)) {
                p.StartInfo.Arguments = args;
            }

            p.OutputDataReceived += new DataReceivedEventHandler(StdOutHandler);
            p.ErrorDataReceived += new DataReceivedEventHandler(StdErrHandler);

            p.Start();
            p.BeginOutputReadLine();
            p.BeginErrorReadLine();

            // I don't know why this sleep is needed but it is :-/
            Thread.Sleep(2000);
            p.PriorityClass = priorityClass;
            p.WaitForExit();

            return p.ExitCode;
        }
    }
}
