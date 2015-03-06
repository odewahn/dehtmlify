Getting Started
===============

In this chapter, we are going to make sure that you have all the prerequisites for doing data science at the command line. The prerequisites fall into two parts: (1) having a proper environment with all the command-line tools that we employ in this book, and (2) understanding the essential concepts that come into play when using the command line.

First, we describe how to install the Data Science Toolbox, which is a virtual environment based on GNU/Linux that contains all the necessary command-line tools. Subsequently, we explain the essential command-line concepts through examples.

By the end of this chapter, you’ll have everything you need in order to continue with the first step of doing data science, namely obtaining data.

Overview
========

In this chapter, you’ll learn:

-   How to set up the Data Science Toolbox

-   Essential concepts and tools necessary to do data science at the command line

Setting Up Your Data Science Toolbox
====================================

 <span id="dstbsetup"></span>   In this book we use many different command-line tools. The distribution of GNU/Linux that we are using, Ubuntu, comes with a whole bunch of command-line tools pre-installed. Moreover, Ubuntu offers many packages that contain other, relevant command-line tools. Installing these packages yourself is not too difficult. However, we also use command-line tools that are not available as packages and require a more manual, and more involved, installation. In order to acquire the necessary command-line tools without having to go through the involved installation process of each, we encourage you to install the Data Science Toolbox.

 If you prefer to run the command-line tools natively rather than inside a virtual machine, then you can install the command-line tools individually. However, be aware that this is a very time-consuming process. [\#\_list\_of\_command\_line\_tools](#_list_of_command_line_tools) lists all the command-line tools used in the book. The installation instructions are for Ubuntu only, so check [the book’s website](http://datascienceatthecommandline.com) for up-to-date information on how to install the command-line tools natively on other operating systems. The scripts and data sets used in the book can be obtained by cloning this book’s [GitHub repository](http://bit.ly/data_science_cl).

The Data Science Toolbox is a virtual environment that allows you to get started doing data science in minutes. The default version comes with commonly used software for data science, including the Python scientific stack and R together with its most popular packages. Additional software and data bundles are easily installed. These bundles can be specific to a certain book, course, or organization. You can read more about the Data Science Toolbox at [its website](http://datasciencetoolbox.org).

There are two ways to set up the Data Science Toolbox: (1) installing it locally using VirtualBox and Vagrant or (2) launching it in the cloud using Amazon Web Services. Both ways result in exactly the same environment. In this chapter, we explain how to set up the Data Science Toolbox for Data Science at the Command Line locally. If you wish to run the Data Science Toolbox in the cloud or if you run into problems, refer to [the book’s website](http://datasciencatthecommandline.com).

The easiest way to install the Data Science Toolbox is on your local machine. Because the local version of the Data Science Toolbox runs on top of VirtualBox and Vagrant, it can be installed on Linux, Mac OS X, and Microsoft Windows.

Step 1: Download and Install VirtualBox
---------------------------------------

 Browse to [the VirtualBox (Oracle, 2014) download page](https://www.virtualbox.org/wiki/Downloads) and download the appropriate binary for your operating system. Open the binary and follow the installation instructions.

Step 2: Download and Install Vagrant
------------------------------------

  Similar to Step 1, browse to [the Vagrant (HashiCorp, 2014) download page](http://www.vagrantup.com/downloads.html) and download the appropriate binary. Open the binary and follow the installation instructions. If you already have Vagrant installed, please make sure that it’s version 1.5 or higher.

Step 3: Download and Start the Data Science Toolbox
---------------------------------------------------

Open a terminal (known as the Command Prompt or PowerShell in Microsoft Windows). Create a directory, e.g., *MyDataScienceToolbox*, and navigate to it by typing:

```
$ mkdir MyDataScienceToolbox
$ cd MyDataScienceToolbox
```

In order to initialize the Data Science Toolbox, run the following command:

```
$ vagrant init data-science-toolbox/data-science-at-the-command-line
```

This creates a file named *Vagrantfile*. This is a configuration file that tells Vagrant how to launch the virtual machine. This file contains a lot of lines that are commented out. A minimal version is shown in [\#vagrantfile\_minimal](#vagrantfile_minimal).

##### Minimal configuration for Vagrant

```
Vagrant.configure(2) do |config|
  config.vm.box = "data-science-toolbox/data-science-at-the-command-line"
end
```

By running the following command, the Data Science Toolbox will be downloaded and booted:

```
$ vagrant up
```

If everything went well, then you now have a Data Science Toolbox running on your local machine.

If you ever see the message `default: Warning: Connection timeout. Retrying...` printed repeatedly, then it may be that the virtual machine is waiting for input. This may happen when the virtual machine has not been properly shut down. In order to find out what’s wrong, add the following lines to *Vagrantfile* before the last `end` statement (also see [\#vagrantfile\_forward](#vagrantfile_forward)):

```
  config.vm.provider "virtualbox" do |vb|
    vb.gui = true
  end
```

This will cause VirtualBox to show a screen. Once the virtual machine has booted and you have identified the problem, you can remove these lines from *Vagrantfile*. The username and password to log in are both `vagrant`. If this doesn’t help, we advise you to check [the book's website](http://datascienceatthecommandline.com), as this website contains an up-to-date list of frequently asked <span class="keep-together">questions</span>.

[\#vagrantfile\_forward](#vagrantfile_forward) shows a slightly more elaborate *Vagrantfile*. You can view more configuration options at [*http://docs.vagrantup.com*](http://docs.vagrantup.com).

##### Configuring Vagrant

```
Vagrant.require_version ">= 1.5.0"                                         
Vagrant.configure(2) do |config|
  config.vm.box = "data-science-toolbox/data-science-at-the-command-line"
  config.vm.network "forwarded_port", guest: 8000, host: 8000              
  config.vm.provider "virtualbox" do |vb|
    vb.gui = true                                                          
    vb.memory = 2048                                                       
    vb.cpus = 2                                                            
  end
end
```

<span id="callout_getting_started_CO1-1">[](#co_getting_started_CO1-1)</span>  
Require at least version 1.5.0 of Vagrant.

<span id="callout_getting_started_CO1-2">[](#co_getting_started_CO1-2)</span>  
Forward port 8000. This is useful if you want to view a figure you created, as we do in [\#chapter7](#chapter7).

<span id="callout_getting_started_CO1-3">[](#co_getting_started_CO1-3)</span>  
Launch a graphical user interface.

<span id="callout_getting_started_CO1-4">[](#co_getting_started_CO1-4)</span>  
Use 2 GB of memory.

<span id="callout_getting_started_CO1-5">[](#co_getting_started_CO1-5)</span>  
Use 2 CPUs.

Step 4: Log In (on Linux and Mac OS X)
--------------------------------------

If you are running Linux, Mac OS X, or some other Unix-like operating system, you can log in to the Data Science Toolbox by running the following command in a <span class="keep-together">terminal</span>:

```
$ vagrant ssh
```

After a few seconds, you will be greeted with the following message:

```
Welcome to the Data Science Toolbox for Data Science at the Command Line

Based on Ubuntu 14.04 LTS (GNU/Linux 3.13.0-24-generic x86_64)

 * Data Science at the Command Line: http://datascienceatthecommandline.com
 * Data Science Toolbox:             http://datasciencetoolbox.org
 * Ubuntu documentation:             http://help.ubuntu.com
Last login: Tue Jul 22 19:33:16 2014 from 10.0.2.2
```

Step 4: Log In (on Microsoft Windows)
-------------------------------------

 If you are running Microsoft Windows, you need to either run Vagrant with a graphical user interface (refer back to Step 2 on how to set that up) or use a third-party application in order to log in to the Data Science Toolbox. For the latter, we <span class="keep-together">recommend</span> PuTTY. Browse to [the PuTTY download page](http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html) and download *putty.exe*. Run PuTTY, and enter the following values:

-   Host Name (or IP address): `127.0.0.1`

-   Port: `2222`

-   Connection type: `SSH`

If you want, you can save these values as a session by clicking the Save button, so that you do not need to enter these values again. Click the Open button and enter `vagrant` for both the username and the password.

Step 5: Shut Down or Start Anew
-------------------------------

 The Data Science Toolbox can be shut down by running the following command from the same directory as you ran `vagrant up`:

```
$ vagrant halt
```

In case you wish to get rid of the Data Science Toolbox and start over, you can type:

```
$ vagrant destroy
```

Then, return to the instructions for Step 3 to set up the Data Science Toolbox again. 

Essential Concepts and Tools
============================

<span id="gnulin"></span> In [\#chapter1](#chapter1), we briefly showed you what the command line is. Now that you have your own Data Science Toolbox, we can really get started. In this section, we discuss several concepts and tools that you will need to know in order to feel comfortable doing data science at the command line. If, up to now, you have been mainly working with graphical user interfaces, this might be quite a change. But don’t worry, we’ll start at the beginning, and very gradually go to more advanced topics.

This section is not a complete course in GNU/Linux. We will only explain the concepts and tools that are relevant for doing data science at the command line. One of the advantages of the Data Science Toolbox is that a lot is already set up. If you wish to know more about GNU/Linux, consult [\#ch2\_further\_reading](#ch2_further_reading) at the end of this chapter.

The Environment
---------------

So you’ve just logged into a brand new environment. Before we do anything, it's worthwhile to get a high-level understanding of this environment. The environment is roughly defined by four layers, which we briefly discuss from the top down:

Command-line tools  
   First and foremost, there are the command-line tools that you work with. We use them by typing their corresponding commands. There are different types of command-line tools, which we will discuss in the next section. Examples of tools are `ls` (Stallman & MacKenzie, 2012), `cat` (Granlund & Stallman, 2012), and `jq` (Dolan, 2014).

Terminal  
 The terminal, which is the second layer, is the application where we type our commands in. If you see the following text:

```
$ seq 3
1
2
3
```

then you would type `seq 3` into your terminal and press **``**. (The command-line tool `seq` (Drepper, 2012) generates a sequence of numbers.) You do not type the dollar sign. It's just there to serve as a prompt and let you know that you can type this command. The text below `seq 3` is the output of the command. In [\#chapter1](#chapter1), we showed you two screenshots of how the default terminal looks in Mac OS X and Ubuntu with various commands and their output. 

Shell  
  The third layer is the shell. Once we have typed in our command and pressed **``**, the terminal sends that command to the shell. The shell is a program that interprets the command. The Data Science Toolbox uses Bash as the shell, but there are many others available. Once you have become a bit more proficient at the command line, you may want to look into a shell called the Z shell. It offers many additional features that can increase your productivity at the command line.

Operating system  
 The fourth layer is the operating system, which is GNU/Linux in our case. Linux is the name of the kernel, which is the heart of the operating system. The kernel is in direct contact with the CPU, disks, and other hardware. The kernel also executes our command-line tools. GNU, which is a recursive acronym for GNU’s Not Unix, refers to a set of basic tools. The Data Science Toolbox is based on a particular Linux distribution called Ubuntu.  

Executing a Command-Line Tool
-----------------------------

 Now that you have an understanding of the environment, it's high time that you try out some commands. Type the following in your terminal (without the dollar sign) and press **``**:

```
$ pwd
/home/vagrant
```

 This is as simple as it gets. You just executed a command that contained a single command-line tool. The command-line tool `pwd` (Meyering, 2012) prints the name of the directory where you currently are. By default, when you log in, this is your home directory. You can view the contents of this directory with `ls` (Stallman & MacKenzie, 2012):

```
$ ls
book
```

 The command-line tool `cd`, which is a Bash builtin, allows you to navigate to a different directory:

```
$ cd book/ch02/
$ cd data
$ pwd
/home/vagrant/book/ch02/data
$ cd ..
$ pwd
/home/vagrant/book/ch02/
```

The part after `cd` specifies to which directory you want to navigate to. Values that come after the command are called command-line arguments or options. The two dots refer to the parent directory. Let’s try a different command:

```
$ head -n 3 data/movies.txt
Matrix
Star Wars
Home Alone
```

 Here we pass three command-line arguments to `head` (MacKenzie & Meyering, 2012). The first one is an option. The second one is a value that belongs to the option. The third one is a filename. This particular command outputs the first three lines of the file *~/book/ch02/data/movies.txt*. 

Sometimes we use commands and pipelines that are too long to fit on the page. In that case, you’ll see something like the following:

```
$ echo 'Hello'\
> ' world' |
> wc
```

The greater-than sign (`>`) is the continuation prompt, which indicates that this line is a continuation of the previous one. A long command can be broken up with either a backslash (`\`) or a pipe symbol (`|`) . Be sure to first match any quotation marks (`"` and `'`). The following command is exactly the same as the previous one:

```
$ echo 'Hello world' | wc
```

Five Types of Command-Line Tools
--------------------------------

<span id="cltbtypes"></span> We employ the term "command-line tool" a lot, but we have not yet explained what we actually mean by it. We use it as an umbrella term for *anything* that can be executed from the command line. Under the hood, each command-line tool is one of the following five types:

-   A binary executable

-   A shell builtin

-   An interpreted script

-   A shell function

-   An alias

It’s good to know the difference between the types. The command-line tools that come pre-installed with the Data Science Toolbox mostly comprise the first two types (binary executable and shell builtin). The other three types (interpreted script, shell function, and alias) allow us to further build up our *own* data science toolbox<span data-type="footnote">Here, we do not refer to the literal Data Science Toolbox we just installed, but to having your own set of tools in a figurative sense.</span> and become more efficient and more productive data scientists.

Binary executable  
 Binary executables are programs in the classical sense. A binary executable is created by compiling source code to machine code. This means that when you open the file in a text editor you cannot read its source code.

Shell builtin  
 Shell builtins are command-line tools provided by the shell, which is Bash in our case. Examples include `cd` and `help`. These cannot be changed. Shell builtins may differ between shells. Like binary executables, they cannot be easily inspected or changed.

Interpreted script  
 An interpreted script is a text file that is executed by a binary executable. Examples include: Python, R, and Bash scripts. One great advantage of an interpreted script is that you can read and change it. [\#script\_fac](#script_fac) shows a script named *~/book/ch02/fac.py*. This script is interpreted by Python not because of the file extension *.py*, but because the first line of the script specifies the binary that should execute it.

##### Python script that computes the factorial of an integer (~/book/ch02/fac.py)

```
#!/usr/bin/env python

def factorial(x):
    result = 1
    for i in xrange(2, x + 1):
        result *= i
    return result

if __name__ == "__main__":
    import sys
    x = int(sys.argv[1])
    print factorial(x)
```

This script computes the factorial of the integer that we pass as a command-line argument. It can be invoked from the command line as follows:

```
$ book/ch02/fac.py 5
120
```

In [\#chapter4](#chapter4), we’ll discuss in great detail how to create reusable command-line tools using interpreted scripts.

Shell function  
   A shell function is a function that is executed by the shell itself; in our case, it is executed by Bash. They provide similar functionality to a Bash script, but they are usually (though not necessarily) smaller than scripts. They also tend to be more personal. The following command defines a function called `fac`, which—just like the interpreted Python script we just looked at—computes the factorial of the integer we pass as a parameter. It does this by generating a list of numbers using `seq`, putting those numbers on one line with `*` as the delimiter using `paste` (Ihnat & MacKenzie, 2012), and passing this equation into `bc` (Nelson, 2006), which evaluates it and outputs the result: 

```
$ fac() { (echo 1; seq $1) | paste -s -d\* | bc; }
$ fac 5
120
```

The file *~/.bashrc*, which is a configuration file for Bash, is a good place to define your shell functions so that they are always available.

Alias  
<span id="alias"></span> Aliases are like macros. If you often find yourself executing a certain command with the same parameters (or a part of it), you can define an alias for this. Aliases are also very useful when you continue to misspell a certain command (see [his GitHub profile](http://bit.ly/mise_chrishwiggins) for a long list of useful aliases). The following commands define two aliases:

```
$ alias l='ls -1 --group-directories-first'
$ alias moer=more
```

Now, if you type the following on the command line, the shell will replace each alias it finds with its value:

```
$ cd ~
$ l
book
```

 Aliases are simpler than shell functions as they don’t allow parameters. The function `fac` could not have been defined using an alias because of the parameter. Still, aliases allow you to save lots of keystrokes. Like shell functions, aliases are often defined in *.bashrc* or *.bash\_aliases* configuration files, which are located in your home directory. To see all aliases currently defined, you simply run `alias` without arguments. Try it, what do you see? 

 In this book, when it comes to creating new command-line tools, we'll focus mostly on the last three types: interpreted scripts, shell functions, and aliases. This is because these can easily be changed. The purpose of a command-line tool is to make your life on the command line easier, and to make you a more productive and more efficient data scientist. You can find out the type of a command-line tool with `type` (which is itself a shell builtin):

```
$ type -a pwd
pwd is a shell builtin
pwd is /bin/pwd
$ type -a cd
cd is a shell builtin
$ type -a fac
fac is a function
fac ()
{
    ( echo 1;
    seq $1 ) | paste -s -d\* | bc
}
$ type -a l
l is aliased to `ls -1 --group-directories-first'
```

As you can see, `type` returns two command-line tools for `pwd`. In that case, the first reported command-line tool is used when you type `pwd`. In the next section, we'll look at how to combine command-line tools.  

Combining Command-Line Tools
----------------------------

  Because most command-line tools adhere to the Unix philosophy, they are designed to do only one thing, and do it really well. For example, the command-line tool `grep` (Meyering, 2012) can filter lines, `wc` (Rubin & MacKenzie, 2012) can count lines, and `sort` (Haertel & Eggert, 2012) can sort lines. The power of the command line comes from its ability to combine these small yet powerful command-line tools. The most common way of combining command-line tools is through a so-called *pipe*. The output from the first tool is passed to the second tool. There are virtually no limits to this.  

 Consider, for example, the command-line tool `seq`, which generates a sequence of numbers. Let's generate a sequence of five numbers:

```
$ seq 5
1
2
3
4
5
```

 The output of a command-line tool is by default passed on to the terminal, which displays it on our screen. We can pipe the ouput of `seq` to a second tool, called `grep`, which can be used to filter lines. Imagine that we only want to see numbers that contain a "3." We can combine `seq` and `grep` as follows:

```
$ seq 30 | grep 3
3
13
23
30
```

 And if we wanted to know how many numbers between 1 and 100 contain a "3", we can use `wc`, which is very good at counting things:

```
$ seq 100 | grep 3 | wc -l
19
```

The `-l` option specifies that `wc` should only output the number of lines. By default it also returns the number of characters and words.

You may start to see that combining command-line tools is a very powerful concept. In the rest of the book, you'll be introduced to many more tools and the functionality they offer when combining them.

Redirecting Input and Output
----------------------------

<span id="cltbredir"></span> We mentioned that, by default, the output of the last command-line tool in the pipeline is outputted to the terminal. You can also save this output to a file. This is called output redirection, and works as follows:

```
$ cd ~/book/ch02
$ seq 10 > data/ten-numbers
```

Here, we save the output of the `seq` tool to a file named *ten-numbers* in the directory *~/book/ch02/data*. If this file does not exist yet, it is created. If this file already did exist, its contents would have been overwritten. You can also append the output to a file with `>>`, meaning the output is put after the original contents:

```
$ echo -n "Hello" > hello-world
$ echo " World" >> hello-world
```

 The tool `echo` simply outputs the value you specify. The `-n` option specifies that `echo` should not output a trailing newline.

Saving the output to a file is useful if you need to store intermediate results (e.g., for continuing with your analysis at a later stage). To use the contents of the file *hello-world* again, we can use `cat` (Granlund & Stallman, 2012), which reads a file and prints it:

```
$ cat hello-world | wc -w
2
```

(Note that the `-w` option indicates `wc` to only count words.) The same result can be achieved with the following notation:

```
$ < hello-world wc -w
2
```

 This way, you are directly passing the file to the standard input of `wc` without running an additional process. If the command-line tool also allows files to be specified as command-line arguments, which many do, you can also do the following for `wc`: 

```
$ wc -w hello-world
2 hello-world
```

Working with Files
------------------

As data scientists, we work with a lot of data, which is often stored in files. It's important to know how to work with files (and the directories they live in) on the command line. Every action that you can do using a graphical user interface, you can do with command-line tools (and much more). In this section, we introduce the most important ones to create, move, copy, rename, and delete files and directories. 

You have already seen how we can create new files by redirecting the output with either `>` or `>>`. In case you need to move a file to a different directory, you can use `mv` (Parker, MacKenzie, & Meyering, 2012): 

```
$ mv hello-world data
```

You can also rename files with `mv`:

```
$ cd data
$ mv hello-world old-file
```

You can also rename or move entire directories. In case you no longer need a file, you delete (or remove) it with `rm` (Rubin, MacKenzie, Stallman, & Meyering, 2012): 

```
$ rm old-file
```

In case you want to remove an entire directory with all its contents, specify the `-r` option, which stands for recursive:

```
$ rm -r ~/book/ch02/data/old
```

In case you want to copy a file, use `cp` (Granlund, MacKenzie, & Meyering, 2012). This is useful for creating backups: 

```
$ cp server.log server.log.bak
```

You can create directories using `mkdir` (MacKenzie, 2012): 

```
$ cd data
$ mkdir logs
```

All of these command-line tools accept the `-v` option, which stands for *verbose*, so that they output what’s going on. All but `mkdir` accept the `-i` option, which stands for *interactive*, and causes the tools to ask you for confirmation.

Using the command-line tools to manage your files can be scary at first, because you have no graphical overview of the filesystem to provide immediate feedback.

Help!
-----

<span id="cltbhelp"></span> As you are finding your way around the command line, it may happen that you need help. Even the most-seasoned Linux users need help at some point. It's impossible to remember all the different command-line tools and their options. Fortunately, the command line offers severals ways to get help.

Perhaps the most important command to get help is perhaps `man` (Eaton & Watson, 2014), which is short for *manual*. It contains information for most command-line tools. Imagine that we forgot the different options to the tool `cat`. You can access its man page using: 

```
$ man cat | head -n 20
CAT(1)                           User Commands                          CAT(1)



NAME
       cat - concatenate files and print on the standard output

SYNOPSIS
       cat [OPTION]... [FILE]...

DESCRIPTION
       Concatenate FILE(s), or standard input, to standard output.

       -A, --show-all
              equivalent to -vET

       -b, --number-nonblank
              number nonempty output lines, overrides -n

       -e     equivalent to -vE
```

   Sometimes you’ll see us use `head`, `fold`, or `cut` at the end of a command. This is only to ensure that the output of the command fits on the page; you don’t have to type these. For example, `head -n 5` only prints the first five lines, `fold` wraps long lines to 80 characters, and `cut -c1-80` trims lines that are longer than 80 characters.

 Not every command-line tool has a man page. For shell builtins, such as `cd`, you need to use the `help` command-line tool:

```
$ help cd | head -n 20
cd: cd [-L|[-P [-e]] [-@]] [dir]
    Change the shell working directory.

    Change the current directory to DIR.  The default DIR is the value of the
    HOME shell variable.

    The variable CDPATH defines the search path for the directory containing
    DIR.  Alternative directory names in CDPATH are separated by a colon (:).
    A null directory name is the same as the current directory.  If DIR begins
    with a slash (/), then CDPATH is not used.

    If the directory is not found, and the shell option 'cdable_vars' is set,
    the word is assumed to be  a variable name.  If that variable has a value,
    its value is used for DIR.

    Options:
        -L  force symbolic links to be followed: resolve symbolic links in
        DIR after processing instances of '..'
        -P  use the physical directory structure without following symbolic
        links: resolve symbolic links in DIR before processing instances
```

 `help` also covers other topics of Bash, in case you are interested (try `help` without command-line arguments for a list of topics).

Newer tools that can be used from the command line often lack a man page as well. In that case, your best bet is to invoke the tool with the `-h` or `--help` option. For example:

```
jq --help

jq - commandline JSON processor [version 1.4]
Usage: jq [options]  [file...]

For a description of the command line options and
how to write jq filters (and why you might want to)
see the jq manpage, or the online documentation at
http://stedolan.github.com/jq
```

Specifying the `--help` option also works for GNU command-line tools, such as `cat`. However, the corresponding man page often provides more information. If after trying these three approaches, you're still stuck, then it's perfectly acceptable to consult the Internet. In [\#\_list\_of\_command\_line\_tools](#_list_of_command_line_tools), there’s a list of all command-line tools used in this book. Besides how each command-line tool can be installed, it also shows how you can get help.  

Further Reading
===============

-   Janssens, J. H. M. (2014). Data Science Toolbox. Retrieved May 10, 2014, from [*http://datasciencetoolbox.org*](http://datasciencetoolbox.org).

-   Oracle. (2014). VirtualBox. Retrieved May 10, 2014, from [*http://virtualbox.org*](http://virtualbox.org).

-   HashiCorp. (2014). Vagrant. Retrieved May 10, 2014, from [*http://vagrantup.com*](http://vagrantup.com).

-   Heddings, L. (2006). Keyboard Shortcuts for Bash. Retrieved May 10, 2014, from [*http://www.howtogeek.com/howto/ubuntu/keyboard-shortcuts-for-bash-command-shell-for-ubuntu-debian-suse-redhat-linux-etc*](http://www.howtogeek.com/howto/ubuntu/keyboard-shortcuts-for-bash-command-shell-for-ubuntu-debian-suse-redhat-linux-etc).

-   Peek, J., Powers, S., O’Reilly, T., & Loukides, M. (2002). [*Unix Power Tools* (3rd Ed.)](http://bit.ly/Unix_Power_Tools_3e). O’Reilly Media.


