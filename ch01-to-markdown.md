<section data-type="chapter" id="chapter2" xmlns="http://www.w3.org/1999/xhtml">

# Getting Started

In this chapter, we are going to make sure that you have all the prerequisites for doing data science at the command line. The prerequisites fall into two parts: (1) having a proper environment with all the command-line tools that we employ in this book, and (2) understanding the essential concepts that come into play when using the command line.

First, we describe how to install the Data Science Toolbox, which is a virtual environment based on GNU/Linux that contains all the necessary command-line tools. Subsequently, we explain the essential command-line concepts through examples.

By the end of this chapter, you’ll have everything you need in order to continue with the first step of doing data science, namely obtaining data.

<section data-type="sect1" id="ch2_overview">

# Overview

In this chapter, you’ll learn:

*   How to set up the Data Science Toolbox

*   Essential concepts and tools necessary to do data science at the command line
</section>

<section data-type="sect1" id="_setting_up_your_data_science_toolbox">

# Setting Up Your Data Science Toolbox

<a contenteditable="false" data-primary="Ubuntu" data-seealso="GNU/Linux" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="Data Science Toolbox" data-secondary="setting up" data-type="indexterm" id="dstbsetup">&nbsp;</a><a contenteditable="false" data-primary="Microsoft Windows" data-see="Data Science Toolbox, setting up" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="Mac OS X" data-see="Data Science Toolbox, setting up" data-type="indexterm">&nbsp;</a>In this book we use many different command-line tools. The distribution of GNU/Linux that we are using, Ubuntu, comes with a whole bunch of command-line tools pre-installed. Moreover, Ubuntu offers many packages that contain other, relevant command-line tools. Installing these packages yourself is not too difficult. However, we also use command-line tools that are not available as packages and require a more manual, and more involved, installation. In order to acquire the necessary command-line tools without having to go through the involved installation process of each, we encourage you to install the Data Science Toolbox.

<div data-type="note"><a contenteditable="false" data-primary="tool" data-seealso="command-line tools" data-type="indexterm">&nbsp;</a>If you prefer to run the command-line tools natively rather than inside a virtual machine, then you can install the command-line tools individually. However, be aware that this is a very time-consuming process. [#_list_of_command_line_tools](#_list_of_command_line_tools) lists all the command-line tools used in the book. The installation instructions are for Ubuntu only, so check [the book’s website](http://datascienceatthecommandline.com) for up-to-date information on how to install the command-line tools natively on other operating systems. The scripts and data sets used in the book can be obtained by cloning this book’s [GitHub repository](http://bit.ly/data_science_cl).</div>

The Data Science Toolbox is a virtual environment that allows you to get started doing data science in minutes. The default version comes with commonly used software for data science, including the Python scientific stack and R together with its most popular packages. Additional software and data bundles are easily installed. These bundles can be specific to a certain book, course, or organization. You can read more about the Data Science Toolbox at [its website](http://datasciencetoolbox.org).

There are two ways to set up the Data Science Toolbox: (1) installing it locally using VirtualBox and Vagrant or (2) launching it in the cloud using Amazon Web Services. Both ways result in exactly the same environment. In this chapter, we explain how to set up the Data Science Toolbox for Data Science at the Command Line locally. If you wish to run the Data Science Toolbox in the cloud or if you run into problems, refer to [the book’s website](http://datasciencatthecommandline.com).

The easiest way to install the Data Science Toolbox is on your local machine. Because the local version of the Data Science Toolbox runs on top of VirtualBox and Vagrant, it can be installed on Linux, Mac OS X, and Microsoft Windows.

<section data-type="sect2" id="_step_1_download_and_install_virtualbox">

## Step 1: Download and Install VirtualBox

<a contenteditable="false" data-primary="VirtualBox" data-type="indexterm">&nbsp;</a>Browse to [the VirtualBox (Oracle, 2014)&nbsp;download page](https://www.virtualbox.org/wiki/Downloads) and download the appropriate binary for your operating system. Open the binary and follow the installation instructions.

</section>

<section data-type="sect2" id="_step_2_download_and_install_vagrant">

## Step 2: Download and Install Vagrant

<a contenteditable="false" data-primary="Vagrant" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="Data Science Toolbox" data-secondary="VirtualBox" data-type="indexterm">&nbsp;</a>Similar to Step 1, browse to [the Vagrant (HashiCorp, 2014) download page](http://www.vagrantup.com/downloads.html) and download the appropriate binary. Open the binary and follow the installation instructions. If you already have Vagrant installed, please make sure that it’s version 1.5 or higher.

</section>

<section data-type="sect2" id="_step_3_download_and_start_the_data_science_toolbox">

## Step 3: Download and Start the Data Science Toolbox

Open a terminal (known as the Command Prompt or PowerShell in Microsoft Windows). Create a directory, e.g., _MyDataScienceToolbox_, and navigate to it by typing:

<pre data-code-language="console" data-type="programlisting">
$ mkdir MyDataScienceToolbox
$ cd MyDataScienceToolbox</pre>

In order to initialize the Data Science Toolbox, run the following command:

<pre data-code-language="console" data-type="programlisting">
$ vagrant init data-science-toolbox/data-science-at-the-command-line</pre>

This creates a file named _Vagrantfile_. This is a configuration file that tells Vagrant how to launch the virtual machine. This file contains a lot of lines that are commented out. A minimal version is shown in [#vagrantfile_minimal](#vagrantfile_minimal).

<div data-type="example" id="vagrantfile_minimal">

##### Minimal configuration for Vagrant

<pre data-code-language="ruby" data-type="programlisting">
Vagrant.configure(2) do |config|
  config.vm.box = "data-science-toolbox/data-science-at-the-command-line"
end</pre>
</div>

By running the following command, the Data Science Toolbox will be downloaded and booted:

<pre data-code-language="console" data-type="programlisting">
$ vagrant up</pre>

If everything went well, then you now have a Data Science Toolbox running on your local machine.

<div data-type="caution">

If you ever see the message `default: Warning: Connection timeout. Retrying...` printed repeatedly, then it may be that the virtual machine is waiting for input. This may happen when the virtual machine has not been properly shut down. In order to find out what’s wrong, add the following lines to _Vagrantfile_ before the last `end` statement (also see [#vagrantfile_forward](#vagrantfile_forward)):

<pre data-code-language="ruby" data-type="programlisting">
  config.vm.provider "virtualbox" do |vb|
    vb.gui = true
  end</pre>

This will cause VirtualBox to show a screen. Once the virtual machine has booted and you have identified the problem, you can remove these lines from _Vagrantfile_. The username and password to log in are both `vagrant`. If this doesn’t help, we advise you to check [the book's website](http://datascienceatthecommandline.com), as this website contains an up-to-date list of frequently asked <span class="keep-together">questions</span>.

</div>

[#vagrantfile_forward](#vagrantfile_forward) shows a slightly more elaborate _Vagrantfile_. You can view more configuration options at [_http://docs.vagrantup.com_](http://docs.vagrantup.com).

<div data-type="example" id="vagrantfile_forward">

##### Configuring Vagrant

<pre data-code-language="ruby" data-type="programlisting">
Vagrant.require_version "&gt;= 1.5.0"                                         [![1](callouts/1.png)](#callout_getting_started_CO1-1)
Vagrant.configure(2) do |config|
  config.vm.box = "data-science-toolbox/data-science-at-the-command-line"
  config.vm.network "forwarded_port", guest: 8000, host: 8000              [![2](callouts/2.png)](#callout_getting_started_CO1-2)
  config.vm.provider "virtualbox" do |vb|
    vb.gui = true                                                          [![3](callouts/3.png)](#callout_getting_started_CO1-3)
    vb.memory = 2048                                                       [![4](callouts/4.png)](#callout_getting_started_CO1-4)
    vb.cpus = 2                                                            [![5](callouts/5.png)](#callout_getting_started_CO1-5)
  end
end</pre>
</div>

<dl class="calloutlist">
	<dt>[![1](callouts/1.png)](#co_getting_started_CO1-1)</dt>
	<dd>

Require at least version 1.5.0 of Vagrant.

	</dd>
	<dt>[![2](callouts/2.png)](#co_getting_started_CO1-2)</dt>
	<dd>

Forward port 8000. This is useful if you want to view a figure you created, as we do in [#chapter7](#chapter7).

	</dd>
	<dt>[![3](callouts/3.png)](#co_getting_started_CO1-3)</dt>
	<dd>

Launch a graphical user interface.

	</dd>
	<dt>[![4](callouts/4.png)](#co_getting_started_CO1-4)</dt>
	<dd>

Use 2 GB of memory.

	</dd>
	<dt>[![5](callouts/5.png)](#co_getting_started_CO1-5)</dt>
	<dd>

Use 2 CPUs.

	</dd>
</dl>
</section>

<section data-type="sect2" id="_step_4_log_in_on_linux_and_mac_os_x">

## Step 4: Log In (on Linux and Mac OS X)

If you are running Linux, Mac OS X, or some other Unix-like operating system, you can log in to the Data Science Toolbox by running the following command in a <span class="keep-together">terminal</span>:

<pre data-code-language="console" data-type="programlisting">
$ vagrant ssh</pre>

After a few seconds, you will be greeted with the following message:

<pre data-code-language="console" data-type="programlisting">
Welcome to the Data Science Toolbox for Data Science at the Command Line

Based on Ubuntu 14.04 LTS (GNU/Linux 3.13.0-24-generic x86_64)

 * Data Science at the Command Line: http://datascienceatthecommandline.com
 * Data Science Toolbox:             http://datasciencetoolbox.org
 * Ubuntu documentation:             http://help.ubuntu.com
Last login: Tue Jul 22 19:33:16 2014 from 10.0.2.2</pre>
</section>

<section data-type="sect2" id="_step_4_log_in_on_microsoft_windows">

## Step 4: Log In (on Microsoft Windows)

<a contenteditable="false" data-primary="Data Science Toolbox" data-secondary="logging in" data-type="indexterm">&nbsp;</a>If you are running Microsoft Windows, you need to either run Vagrant with a graphical user interface (refer back to Step 2 on how to set that up) or use a third-party application in order to log in to the Data Science Toolbox. For the latter, we <span class="keep-together">recommend</span> PuTTY. Browse to [the PuTTY download page](http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html) and download _putty.exe_. Run PuTTY, and enter the following values:

*   Host Name (or IP address): `127.0.0.1`

*   Port: `2222`

*   Connection type: `SSH`

If you want, you can save these values as a session by clicking the Save button, so that you do not need to enter these values again. Click the Open button and enter `vagrant` for both the username and the password.

</section>

<section data-type="sect2" id="_step_5_shutdown_or_start_anew">

## Step 5: Shut Down or Start Anew

<a contenteditable="false" data-primary="Data Science Toolbox" data-secondary="starting anew" data-type="indexterm">&nbsp;</a>The Data Science Toolbox can be shut down by running the following command from the same directory as you ran `vagrant up`:

<pre data-code-language="console" data-type="programlisting">
$ vagrant halt</pre>

In case you wish to get rid of the Data Science Toolbox and start over, you can type:

<pre data-code-language="console" data-type="programlisting">
$ vagrant destroy</pre>

Then, return to the instructions for Step 3 to set up the Data Science Toolbox again.<a contenteditable="false" data-primary="Data Science Toolbox" data-secondary="setting up" data-startref="dstbsetup" data-type="indexterm">&nbsp;</a>

</section>
</section>

<section data-type="sect1" id="_essential_gnu_linux_concepts">

# Essential Concepts and Tools

<a contenteditable="false" data-primary="GNU/Linux" data-type="indexterm" id="gnulin">&nbsp;</a>In [#chapter1](#chapter1), we briefly showed you what the command line is. Now that you have your own Data Science Toolbox, we can really get started. In this section, we discuss several concepts and tools that you will need to know in order to feel comfortable doing data science at the command line. If, up to now, you have been mainly working with graphical user interfaces, this might be quite a change. But don’t worry, we’ll start at the beginning, and very gradually go to more advanced topics.

<div data-type="note">This section is not a complete course in GNU/Linux. We will only explain the concepts and tools that are relevant for doing data science at the command line. One of the advantages of the Data Science Toolbox is that a lot is already set up. If you wish to know more about GNU/Linux, consult [#ch2_further_reading](#ch2_further_reading) at the end of this chapter.</div>

<section data-type="sect2" id="_the_environment">

## The Environment

So you’ve just logged into a brand new environment. Before we do anything, it's worthwhile to get a high-level understanding of this environment. The environment is roughly defined by four layers, which we briefly discuss from the top down:

<dl>
	<dt>Command-line tools</dt>
	<dd>

<a contenteditable="false" data-primary="command-line tools" data-seealso="Data Science Toolbox" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="tool" data-type="indexterm">&nbsp;</a>&nbsp;First and foremost, there are the command-line tools that you work with. We use them by typing their corresponding commands. There are different types of command-line tools, which we will discuss in the next section. Examples of tools are `ls` (Stallman &amp; MacKenzie, 2012), `cat` (Granlund &amp; Stallman, 2012), and `jq` (Dolan, 2014).

	</dd>
	<dt>Terminal</dt>
	<dd>

<a contenteditable="false" data-primary="terminal" data-type="indexterm">&nbsp;</a>The terminal, which is the second layer, is the application where we type our commands in. If you see the following text:

	<pre data-code-language="console" data-type="programlisting">
$ seq 3
1
2
3</pre>

then you would type `seq 3` into your terminal and press **`<Enter>`**. (The command-line tool `seq` (Drepper, 2012) generates a sequence of numbers.) You do not type the dollar sign. It's just there to serve as a prompt and let you know that you can type this command. The text below `seq 3` is the output of the command. In [#chapter1](#chapter1), we showed you two screenshots of how the default terminal looks in Mac OS X and Ubuntu with various commands and their output.<a contenteditable="false" data-primary="seq" data-type="indexterm">&nbsp;</a>

	</dd>
	<dt>Shell</dt>
	<dd>

<a contenteditable="false" data-primary="Bash" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="shell" data-type="indexterm">&nbsp;</a>The third layer is the shell. Once we have typed in our command and pressed **`<Enter>`**, the terminal sends that command to the shell. The shell is a program that interprets the command. The Data Science Toolbox uses Bash as the shell, but there are many others available. Once you have become a bit more proficient at the command line, you may want to look into a shell called the Z shell. It offers many additional features that can increase your productivity at the command line.

	</dd>
	<dt>Operating system</dt>
	<dd>

&nbsp;The fourth layer is the operating system, which is GNU/Linux in our case. Linux is the name of the kernel, which is the heart of the operating system. The kernel is in direct contact with the CPU, disks, and other hardware. The kernel also executes our command-line tools. GNU, which is a recursive acronym for GNU’s Not Unix, refers to a set of basic tools. The Data Science Toolbox is based on a particular Linux distribution called Ubuntu.&nbsp;&nbsp;

	</dd>
</dl>
</section>

<section data-type="sect2" id="_executing_a_command_line_tool">

## Executing a Command-Line Tool

<a contenteditable="false" data-primary="command-line tools" data-secondary="executing" data-type="indexterm">&nbsp;</a>Now that you have an understanding of the environment, it's high time that you try out some commands. Type the following in your terminal (without the dollar sign) and press **`<Enter>`**:

<pre data-code-language="console" data-type="programlisting">
$ pwd
/home/vagrant</pre>

<a contenteditable="false" data-primary="pwd" data-type="indexterm">&nbsp;</a>This is as simple as it gets. You just executed a command that contained a single command-line tool. The command-line tool `pwd` (Meyering, 2012) prints the name of the directory where you currently are. By default, when you log in, this is your home directory. You can view the contents of this directory with `ls` (Stallman &amp; MacKenzie, 2012):

<pre data-code-language="console" data-type="programlisting">
$ ls
book</pre>

<a contenteditable="false" data-primary="cd" data-type="indexterm">&nbsp;</a>The command-line tool `cd`, which is a Bash builtin, allows you to navigate to a different directory:

<pre data-code-language="console" data-type="programlisting">
$ cd book/ch02/
$ cd data
$ pwd
/home/vagrant/book/ch02/data
$ cd ..
$ pwd
/home/vagrant/book/ch02/</pre>

The part after `cd` specifies to which directory you want to navigate to. Values that come after the command are called command-line arguments or options. The two dots refer to the parent directory. Let’s try a different command:

<pre data-code-language="console" data-type="programlisting">
$ head -n 3 data/movies.txt
Matrix
Star Wars
Home Alone</pre>

<a contenteditable="false" data-primary="head" data-type="indexterm">&nbsp;</a>Here we pass three command-line arguments to `head` (MacKenzie &amp; Meyering, 2012). The first one is an option. The second one is a value that belongs to the option. The third one is a filename. This particular command outputs the first three lines of the file _~/book/ch02/data/movies.txt_.&nbsp;

Sometimes we use commands and pipelines that are too long to fit on the page. In that case, you’ll see something like the following:

<pre data-code-language="console" data-type="programlisting">
$ echo 'Hello'\
&gt; ' world' |
&gt; wc</pre>

The greater-than sign (`>`) is the continuation prompt, which indicates that this line is a continuation of the previous one. A long command can be broken up with either a backslash (`\`) or a pipe symbol (`|`) . Be sure to first match any quotation marks (`"`&nbsp;and&nbsp;`'`). The following command is exactly the same as the previous one:

<pre data-code-language="console" data-type="programlisting">
$ echo 'Hello world' | wc</pre>
</section>

<section data-type="sect2" id="_five_types_of_command_line_tools">

## Five Types of Command-Line Tools

<a contenteditable="false" data-primary="command-line tools" data-secondary="types" data-type="indexterm" id="cltbtypes">&nbsp;</a>We employ the term "command-line tool" a lot, but we have not yet explained what we actually mean by it. We use it as an umbrella term for _anything_ that can be executed from the command line. Under the hood, each command-line tool is one of the following five types:

*   A binary executable

*   A shell builtin

*   An interpreted script

*   A shell function

*   An alias

It’s good to know the difference between the types. The command-line tools that come pre-installed with the Data Science Toolbox mostly comprise the first two types (binary executable and shell builtin). The other three types (interpreted script, shell function, and alias) allow us to further build up our _own_ data science toolbox<span data-type="footnote">Here, we do not refer to the literal Data Science Toolbox we just installed, but to having your own set of tools in a figurative sense.</span> and become more efficient and more productive data scientists.

<dl>
	<dt>Binary executable</dt>
	<dd>

<a contenteditable="false" data-primary="binary executables" data-type="indexterm">&nbsp;</a>Binary executables are programs in the classical sense. A binary executable is created by compiling source code to machine code. This means that when you open the file in a text editor you cannot read its source code.

	</dd>
	<dt>Shell builtin</dt>
	<dd>

<a contenteditable="false" data-primary="shell builtins" data-type="indexterm">&nbsp;</a>Shell builtins are command-line tools provided by the shell, which is Bash in our case. Examples include `cd` and `help`. These cannot be changed. Shell builtins may differ between shells. Like binary executables, they cannot be easily inspected or changed.

	</dd>
	<dt>Interpreted script</dt>
	<dd>

<a contenteditable="false" data-primary="interpreted scripts" data-type="indexterm">&nbsp;</a>An interpreted script is a text file that is executed by a binary executable. Examples include: Python, R, and Bash scripts. One great advantage of an interpreted script is that you can read and change it. [#script_fac](#script_fac) shows a script named _~/book/ch02/fac.py_. This script is interpreted by Python not because of the file extension _.py_, but because the first line of the script specifies the binary that should execute it.

	<div data-type="example" id="script_fac">

##### Python script that computes the factorial of an integer (~/book/ch02/fac.py)

	<pre data-code-language="python" data-type="programlisting">
#!/usr/bin/env python

def factorial(x):
    result = 1
    for i in xrange(2, x + 1):
        result *= i
    return result

if __name__ == "__main__":
    import sys
    x = int(sys.argv[1])
    print factorial(x)</pre>
	</div>

This script computes the factorial of the integer that we pass as a command-line argument. It can be invoked from the command line as follows:

	<pre data-code-language="console" data-type="programlisting">
$ book/ch02/fac.py 5
120</pre>

In [#chapter4](#chapter4), we’ll discuss in great detail how to create reusable command-line tools using interpreted scripts.

	</dd>
	<dt>Shell function</dt>
	<dd>

<a contenteditable="false" data-primary="fac" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="Bash" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="shell functions" data-type="indexterm">&nbsp;</a>A shell function is a function that is executed by the shell itself; in our case, it is executed by Bash. They provide similar functionality to a Bash script, but they are usually (though not necessarily) smaller than scripts. They also tend to be more personal. The following command defines a function called `fac`, which—just like the interpreted Python script we just looked at—computes the factorial of the integer we pass as a parameter. It does this by generating a list of numbers using `seq`, putting those numbers on one line with `*` as the delimiter using `paste` (Ihnat &amp; MacKenzie, 2012), and passing this equation into `bc` (Nelson, 2006), which evaluates it and outputs the result:<a contenteditable="false" data-primary="paste" data-type="indexterm">&nbsp;</a>

	<pre data-code-language="console" data-type="programlisting">
$ fac() { (echo 1; seq $1) | paste -s -d\* | bc; }
$ fac 5
120</pre>

The file _~/.bashrc_, which is a configuration file for Bash, is a good place to define your shell functions so that they are always available.

	</dd>
	<dt>Alias</dt>
	<dd>

<a contenteditable="false" data-primary="alias" data-type="indexterm" id="alias">&nbsp;</a>Aliases are like macros. If you often find yourself executing a certain command with the same parameters (or a part of it), you can define an alias for this. Aliases are also very useful when you continue to misspell a certain command (see [his GitHub profile](http://bit.ly/mise_chrishwiggins) for a long list of useful aliases). The following commands define two aliases:

	<pre data-code-language="console" data-type="programlisting">
$ alias l='ls -1 --group-directories-first'
$ alias moer=more</pre>

Now, if you type the following on the command line, the shell will replace each alias it finds with its value:

	<pre data-code-language="console" data-type="programlisting">
$ cd ~
$ l
book</pre>

&nbsp;Aliases are simpler than shell functions as they don’t allow parameters. The function `fac` could not have been defined using an alias because of the parameter. Still, aliases allow you to save lots of keystrokes. Like shell functions, aliases are often defined in _.bashrc_ or _.bash_aliases_ configuration files, which are located in your home directory. To see all aliases currently defined, you simply run `alias` without arguments. Try it, what do you see?<a contenteditable="false" data-primary="alias" data-startref="alias" data-type="indexterm">&nbsp;</a>

	</dd>
</dl>

<a contenteditable="false" data-primary="type" data-type="indexterm">&nbsp;</a>In this book, when it comes to creating new command-line tools, we'll focus mostly on the last three types: interpreted scripts, shell functions, and aliases. This is because these can easily be changed. The purpose of a command-line tool is to make your life on the command line easier, and to make you a more productive and more efficient data scientist. You can find out the type of a command-line tool with `type` (which is itself a shell builtin):

<pre data-code-language="console" data-type="programlisting">
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
l is aliased to `ls -1 --group-directories-first'</pre>

As you can see, `type` returns two command-line tools for `pwd`. In that case, the first reported command-line tool is used when you type `pwd`. In the next section, we'll look at how to combine command-line tools.<a contenteditable="false" data-primary="pwd" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="command-line tools" data-secondary="types" data-startref="cltbtypes" data-type="indexterm">&nbsp;</a>

</section>

<section data-type="sect2" id="_combining_command_line_tools">

## Combining Command-Line Tools

<a contenteditable="false" data-primary="command-line tools" data-secondary="combining" data-seealso="pipe" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="Unix" data-secondary="philosophy" data-type="indexterm">&nbsp;</a>Because most command-line tools adhere to the Unix philosophy, they are designed to do only one thing, and do it really well. For example, the command-line tool `grep` (Meyering, 2012) can filter lines, `wc` (Rubin &amp; MacKenzie, 2012) can count lines, and `sort` (Haertel &amp; Eggert, 2012) can sort lines. The power of the command line comes from its ability to combine these small yet powerful command-line tools. The most common way of combining command-line tools is through a so-called _pipe_. The output from the first tool is passed to the second tool. There are virtually no limits to this.<a contenteditable="false" data-primary="pipe" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="Unix" data-secondary="pipe" data-type="indexterm">&nbsp;</a>

<a contenteditable="false" data-primary="seq" data-type="indexterm">&nbsp;</a>Consider, for example, the command-line tool `seq`, which generates a sequence of numbers. Let's generate a sequence of five numbers:

<pre data-code-language="console" data-type="programlisting">
$ seq 5
1
2
3
4
5</pre>

<a contenteditable="false" data-primary="grep" data-type="indexterm">&nbsp;</a>The output of a command-line tool is by default passed on to the terminal, which displays it on our screen. We can pipe the ouput of `seq` to a second tool, called `grep`, which can be used to filter lines. Imagine that we only want to see numbers that contain a "3." We can combine `seq` and `grep` as follows:

<pre data-code-language="console" data-type="programlisting">
$ seq 30 | grep 3
3
13
23
30</pre>

<a contenteditable="false" data-primary="wc" data-type="indexterm">&nbsp;</a>And if we wanted to know how many numbers between 1 and 100 contain a "3", we can use `wc`, which is very good at counting things:

<pre data-code-language="console" data-type="programlisting">
$ seq 100 | grep 3 | wc -l
19</pre>

The `-l` option specifies that `wc` should only output the number of lines. By default it also returns the number of characters and words.

You may start to see that combining command-line tools is a very powerful concept. In the rest of the book, you'll be introduced to many more tools and the functionality they offer when combining them.

</section>

<section data-type="sect2" id="_redirecting_input_and_output">

## Redirecting Input and Output

<a contenteditable="false" data-primary="command-line tools" data-secondary="redirecting input and output" data-type="indexterm" id="cltbredir">&nbsp;</a>We mentioned that, by default, the output of the last command-line tool in the pipeline is outputted to the terminal. You can also save this output to a file. This is called output redirection, and works as follows:

<pre data-code-language="console" data-type="programlisting">
$ cd ~/book/ch02
$ seq 10 &gt; data/ten-numbers</pre>

Here, we save the output of the `seq` tool to a file named _ten-numbers_ in the directory _~/book/ch02/data_. If this file does not exist yet, it is created. If this file already did exist, its contents would have been overwritten. You can also append the output to a file with `>>`, meaning the output is put after the original contents:

<pre data-code-language="console" data-type="programlisting">
$ echo -n "Hello" &gt; hello-world
$ echo " World" &gt;&gt; hello-world</pre>

<a contenteditable="false" data-primary="echo" data-type="indexterm">&nbsp;</a>The tool `echo` simply outputs the value you specify. The `-n` option specifies that `echo` should not output a trailing newline.

Saving the output to a file is useful if you need to store intermediate results (e.g., for continuing with your analysis at a later stage). To use the contents of the file _hello-world_ again, we can use `cat` (Granlund &amp; Stallman, 2012), which reads a file and prints it:

<pre data-code-language="console" data-type="programlisting">
$ cat hello-world | wc -w
2</pre>

(Note that the `-w` option indicates `wc` to only count words.) The same result can be achieved with the following notation:

<pre data-code-language="console" data-type="programlisting">
$ &lt; hello-world wc -w
2</pre>

<a contenteditable="false" data-primary="command-line tools" data-secondary="redirecting input and output" data-startref="cltbredir" data-type="indexterm">&nbsp;</a>This way, you are directly passing the file to the standard input of `wc` without running an additional process. If the command-line tool also allows files to be specified as command-line arguments, which many do, you can also do the following for `wc`:&nbsp;

<pre data-code-language="console" data-type="programlisting">
$ wc -w hello-world
2 hello-world</pre>
</section>

<section data-type="sect2" id="_working_with_files">

## Working with Files

As data scientists, we work with a lot of data, which is often stored in files. It's important to know how to work with files (and the directories they live in) on the command line. Every action that you can do using a graphical user interface, you can do with command-line tools (and much more). In this section, we introduce the most important ones to create, move, copy, rename, and delete files and directories.<a contenteditable="false" data-primary="command-line tools" data-secondary="files" data-type="indexterm">&nbsp;</a>

You have already seen how we can create new files by redirecting the output with either `>` or `>>`. In case you need to move a file to a different directory, you can use `mv` (Parker, MacKenzie, &amp; Meyering, 2012):<a contenteditable="false" data-primary="mv" data-type="indexterm">&nbsp;</a>

<pre data-code-language="console" data-type="programlisting">
$ mv hello-world data</pre>

You can also rename files with `mv`:

<pre data-code-language="console" data-type="programlisting">
$ cd data
$ mv hello-world old-file</pre>

You can also rename or move entire directories. In case you no longer need a file, you delete (or remove) it with `rm` (Rubin, MacKenzie, Stallman, &amp; Meyering, 2012):<a contenteditable="false" data-primary="rm" data-type="indexterm">&nbsp;</a>

<pre data-code-language="console" data-type="programlisting">
$ rm old-file</pre>

In case you want to remove an entire directory with all its contents, specify the `-r` option, which stands for recursive:

<pre data-code-language="console" data-type="programlisting">
$ rm -r ~/book/ch02/data/old</pre>

In case you want to copy a file, use `cp` (Granlund, MacKenzie, &amp; Meyering, 2012). This is useful for creating backups:<a contenteditable="false" data-primary="cp" data-type="indexterm">&nbsp;</a>

<pre data-code-language="console" data-type="programlisting">
$ cp server.log server.log.bak</pre>

You can create directories using `mkdir` (MacKenzie, 2012):<a contenteditable="false" data-primary="mkdir" data-type="indexterm">&nbsp;</a>

<pre data-code-language="console" data-type="programlisting">
$ cd data
$ mkdir logs</pre>

All of these command-line tools accept the `-v` option, which stands for _verbose_, so that they output what’s going on. All but `mkdir` accept the `-i` option, which stands for _interactive_, and causes the tools to ask you for confirmation.

Using the command-line tools to manage your files can be scary at first, because you have no graphical overview of the filesystem to provide immediate feedback.

</section>

<section data-type="sect2" id="_help">

## Help!

<a contenteditable="false" data-primary="command-line tools" data-secondary="getting help" data-type="indexterm" id="cltbhelp">&nbsp;</a>As you are finding your way around the command line, it may happen that you need help. Even the most-seasoned Linux users need help at some point. It's impossible to remember all the different command-line tools and their options. Fortunately, the command line offers severals ways to get help.

Perhaps the most important command to get help is perhaps `man` (Eaton &amp; Watson, 2014), which is short for _manual_. It contains information for most command-line tools. Imagine that we forgot the different options to the tool `cat`. You can access its man page using:<a contenteditable="false" data-primary="man" data-type="indexterm">&nbsp;</a>

<pre data-code-language="console" data-type="programlisting">
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

       -e     equivalent to -vE</pre>

<div data-type="tip"><a contenteditable="false" data-primary="head" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="fold" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="cut" data-type="indexterm">&nbsp;</a>Sometimes you’ll see us use `head`, `fold`, or `cut` at the end of a command. This is only to ensure that the output of the command fits on the page; you don’t have to type these. For example, `head -n 5` only prints the first five lines, `fold` wraps long lines to 80 characters, and `cut -c1-80` trims lines that are longer than 80 characters.</div>

<a contenteditable="false" data-primary="help" data-type="indexterm">&nbsp;</a>Not every command-line tool has a man page. For shell builtins, such as `cd`, you need to use the `help` command-line tool:

<pre data-code-language="console" data-type="programlisting">
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
        -L	force symbolic links to be followed: resolve symbolic links in
    	DIR after processing instances of '..'
        -P	use the physical directory structure without following symbolic
    	links: resolve symbolic links in DIR before processing instances</pre>

<a contenteditable="false" data-primary="type" data-type="indexterm">&nbsp;</a>`help` also covers other topics of Bash, in case you are interested (try `help` without command-line arguments for a list of topics).

Newer tools that can be used from the command line often lack a man page as well. In that case, your best bet is to invoke the tool with the `-h` or `--help` option. For example:

<pre data-code-language="console" data-type="programlisting">
jq --help

jq - commandline JSON processor [version 1.4]
Usage: jq [options] &lt;jq filter&gt; [file...]

For a description of the command line options and
how to write jq filters (and why you might want to)
see the jq manpage, or the online documentation at
http://stedolan.github.com/jq</pre>

Specifying the `--help` option also works for GNU command-line tools, such as `cat`. However, the corresponding man page often provides more information. If after trying these three approaches, you're still stuck, then it's perfectly acceptable to consult the Internet. In [#_list_of_command_line_tools](#_list_of_command_line_tools), there’s a list of all command-line tools used in this book. Besides how each command-line tool can be installed, it also shows how you can get help.<a contenteditable="false" data-primary="GNU/Linux" data-startref="gnulin" data-type="indexterm">&nbsp;</a><a contenteditable="false" data-primary="command-line tools" data-secondary="getting help" data-startref="cltbhelp" data-type="indexterm">&nbsp;</a>

</section>
</section>

<section data-type="sect1" id="ch2_further_reading">

# Further Reading

*   Janssens, J. H. M. (2014). Data Science Toolbox. Retrieved May 10, 2014, from [_http://datasciencetoolbox.org_](http://datasciencetoolbox.org).

*   Oracle. (2014). VirtualBox. Retrieved May 10, 2014, from [_http://virtualbox.org_](http://virtualbox.org).

*   HashiCorp. (2014). Vagrant. Retrieved May 10, 2014, from [_http://vagrantup.com_](http://vagrantup.com).

*   Heddings, L. (2006). Keyboard Shortcuts for Bash. Retrieved May 10, 2014, from [_http://www.howtogeek.com/howto/ubuntu/keyboard-shortcuts-for-bash-command-shell-for-ubuntu-debian-suse-redhat-linux-etc_](http://www.howtogeek.com/howto/ubuntu/keyboard-shortcuts-for-bash-command-shell-for-ubuntu-debian-suse-redhat-linux-etc).

*   Peek, J., Powers, S., O’Reilly, T., &amp; Loukides, M. (2002). [_Unix Power Tools_ (3rd Ed.)](http://bit.ly/Unix_Power_Tools_3e). O’Reilly Media.
</section>
</section>