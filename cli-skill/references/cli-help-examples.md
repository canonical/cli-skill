# **Canonical CLI Help**

# **Command line tool feedback**

Interacting with a command-line interface is similar to a dialog in that the user issues commands, then has a chance to reflect on the feedback before issuing another command. It requires the user to be precise, and follow the intended structure when issuing a command. As that can be complex, and as users cannot be counted upon to care enough about a command to internalize its logic, it is vital that every tool strives to create clarity of this interaction: what commands are available, what will be the consequences of actions, how can the current state be queried.

The first step is to provide help. This help needs to be concise enough to be digestible and allow a quick grasp of the shape and scope of a command, but must also be extensive enough that a new or infrequent user is given enough information to understand the help output, so that every answer becomes “helpful”.    

## **Help Structure**

###### *Status Quo (derived from observation)*

Every command line tool should provide a help function:

| `tool help tool help --all tool help <COMMAND>` |

The output of the help function should be text to stdout following a standard structure:

* Description: provides a high level description of the purpose of the tool.  
* Usage: list of parameters, commands and options as used on the command line.   
* Command list: list of commands; long lists grouped by topic.  
* Instructions: how to get more information.

```
$ foo help
foo creates and manages foos running in a remote bar.

Usage:
  foo <command> [<options>...] argument

Commonly used commands can be classified as follows:

          Lifecycle: start, stop
   Image management: pull
     Bar management: attach, disband, clean

For more information on a command, run 'foo help <command>'.
For a summary of all commands, run 'foo help --all'
```

Every command line tool should also provide a help function for commands: 

* Usage: list of parameters, commands and options as used on the command line.   
* Description: provides a detailed description of the command, describing default behaviour and implications of using flags.  
* Flag list: list of flags.

```
$ foo help start
Usage:
  foo start [<options>...] <bar>...

The start command creates foos running in a remote bar.

[start command options]
      --color=[auto|never|always]          Use a little bit of colour to highlight some things. (default: auto)
      --unicode=[auto|never|always]        Use a little bit of Unicode to improve legibility. (default: auto)
      --no-wait                            Do not wait for the operation to finish but just print the change id.
      --channel=                           Use this channel instead of stable
      --edge                               Install from the edge channel
      --beta                               Install from the beta channel

```

##### Proposal

Every command line tool should provide a help function; as the user is looking for help, it is important the syntax is forgiving, and it is vital tools understand all variants (help both as a command and a flag):

```
tool help
tool --help
tool -h

tool help <COMMAND>
tool <COMMAND> --help
```

The output of the help function should be text to stdout following a standard structure that aligns well with other tools that the user might be using; we are basing this recommendation on universal standards, and as our tools are usually complex with more than a handful of commands, adding a focus on meaningful grouping of commands:

* Usage: list of parameters, commands and options as used on the command line.   
* Summary: provides a high level description of the purpose of the tool.  
* (Groups of) Commands: one-line descriptions for (most important) commands. Long lists grouped by topic / object.  
* Instructions on how to get more information  
* (Optional) suggestions of proper usage or how to initialise (if config not found)

<pre>
$ foo help
<b>foo 1.21</b>
<span style="color: gray; font-weight: bold">Usage:</span>
  foo <command> [<options>...] argument

  foo creates and manages foos running in a remote bar.

<span style="color: gray; font-weight: bold">Global options:</span>
  --help       Show this help message and exit
  --verbose    Show debug information and be more verbose
  --quiet      Only show warnings and errors, not progress
  --verbosity  Set verbosity: 'quiet', 'brief', 'verbose', 'debug' or 'trace'
  --version    Show the application version and exit

<span style="color: gray; font-weight: bold">Lifecycle:</span>
  start   Start a new foo
  stop    Stop a running foo

<span style="color: gray; font-weight: bold">Image management:</span>
  pull    Pull an image from the central repository

<span style="color: gray; font-weight: bold">Bar management:</span>
  attach  Attach a new bar to manage

For more information on a command, use 'TOOL help <command>'
For a summary of all commands, use 'TOOL help --all'
For more help read the docs: https://tool.readthedocs.org

<span style="color: gray; font-weight: bold">Examples:</span>
  foo attach bar --transient     attaches a new bar in transient mode

foo is unable to find a config file, run <b>foo init</b> to create a configuration file in ~/.local/foo and install command completions for your current shell. 
</pre>

<pre>
$ rockcraft help
<b>rockcraft 1.7.0.post26+g8ea35b0</b>
<span style="color: gray; font-weight: bold">Usage:</span>
    rockcraft <command> [<options>...]

    Rockcraft is a tool to create rocks – a new generation of secure, stable 
    and OCI-compliant container images (docker images), based on Ubuntu.

<span style="color: gray; font-weight: bold">Global options:</span>
  --help       Show this help message and exit
  --quiet      Only show warnings and errors, not progress
  --verbose    More verbose output, repeat to increase detail
  --version    Show the application version and exit

<span style="color: gray; font-weight: bold">General commands:</span>
  init         Create a minimal rockcraft.yaml in current directory

<span style="color: gray; font-weight: bold">Lifecycle commands (rockcraft help lifecycle for more details):</span>
  clean        Remove assets for a part
  pull         Download or retrieve artefacts defined for a part
  build        Build artefacts defined for a part
  overlay      Combine overlay layer for parts (optional)
  stage        Stage build artefacts into a shared area
  prime        Prepare payload to be packed, adding metadata files
  pack         Create the final artefact

<span style="color: gray; font-weight: bold">Other topics:</span>
     <u>Account</u>   login, logout, whoami, export-credentials
  <u>Extensions</u>   expand-extensions, extensions

For more information about a topic, run 'rockcraft help <topic>'.
For more information about a command, run 'rockcraft help <command>' or browse the <u>rockcraft reference documentation</u>.

<span style="color: gray; font-weight: bold">Suggestions:</span>
<span style="color: red">!</span> No config file found in current directory, consider running 'rockcraft init'
</pre>

<pre>
$ snap --help
<b>snap 2.67</b>
<span style="color: gray; font-weight: bold">Usage:</span>
    snap <command> [<options>...]

    The snap command lets you install, configure, refresh and remove snaps.
    Snaps are packages that work across many different Linux distributions,
    enabling secure delivery and operation of the latest apps and utilities.

<span style="color: gray; font-weight: bold">Global options:</span>
  --help       Show this help message and exit
  --quiet      Only show warnings and errors, not progress
  --verbose    More verbose output, repeat to increase detail
  --version    Show the application version and exit

<span style="color: gray; font-weight: bold">Basic commands:</span>
  find             Find packages to install
  install          Install snaps on the system
  refresh          Refresh snaps in the system
  remove           Remove snaps from the system
  info             Show detailed information about snaps
  list             List installed snaps

<span style="color: gray; font-weight: bold">Other topics:</span>
      <u>Management</u>   components, revert, switch, enable, disable, create-cohort
     <u>Permissions</u>   connections, interface, connect, disconnect
         <u>History</u>   changes, tasks, abort, watch
         <u>Daemons</u>   services, start, stop, restart, logs
   <u>Configuration</u>   get, set, unset, wait
         <u>Aliases</u>   alias, aliases, unalias, prefer
         <u>Account</u>   login, logout, whoami
       <u>Snapshots</u>   saved, save, check-snapshot, restore, forget
          <u>Device</u>   model, remodel, reboot, recovery
          <u>Quotas</u>   set-quota, remove-quota, quotas, quota
 Validation <u>Sets</u>   validate
        <u>Warnings</u>   warnings, okay
      <u>Assertions</u>   known, ack

For more information about a command, run 'snap help <command>'.
For more information about a topic, run 'snap help <topic>', eg 'help history'
For a short summary of all commands, run 'snap help --all'.
</pre>

<pre>
$ snap help install
<b>snap 2.67</b>
<span style="color: gray; font-weight: bold">Usage:</span>
  snap install [&lt;option&gt;...] &lt;snap&gt;...

The install command installs the named snaps on the system.

With no further options, the snaps are installed tracking the stable channel, with strict security 
confinement. All available channels of a snap are listed in its 'snap info' output.


<span style="color: gray; font-weight: bold">Global options:</span>
  --help              Show this help message and exit
  --color=[n|y]       Use a little bit of colour to highlight some things. (default: auto)
  --unicode=[n|y]     Use a little bit of Unicode to improve legibility. (default: auto)
  --no-wait           Do not wait for the operation to finish but just print the change id.

<span style="color: gray; font-weight: bold">Channel and revision options:</span>
  --stable                Install from the stable channel (default)
  --candidate             Install from the candidate channel
  --beta                  Install from the beta channel
  --edge                  Install from the edge channel
  --channel=<channel>     Use the specified channel instead of stable
  --revision=             Install the given revision of a snap

<span style="color: gray; font-weight: bold">Installation mode options:</span>
  --classic               Put snap in classic mode and disable security confinement
  --jailmode              Put snap in enforced confinement mode
  --devmode               Put snap in development mode and disable security confinement
  --dangerous             Install the given snap file even if there are no pre-acknowledged signatures
                          for it, meaning it was not verified and could be dangerous (implied by --devmode)
  --unaliased             Install the given snap without enabling its automatic aliases
  --prefer                Enable all aliases of the given snap in preference to conflicting aliases
  --parallel=<suffix>     Install multiple instances of the same snap, appending _suffix to the snap's name

<span style="color: gray; font-weight: bold">Other options:</span>
  --name=<name>           Install the snap file under the given instance name when installing from a file
  --cohort=<cohort>       Install the snap in the given cohort
  --ignore-validation     Ignore validation by other snaps blocking the installation
  --transaction=<scope>   Have one transaction [per-snap] (default) or one for [all-snaps]
  --quota-group=<group>   Add the snap to a quota group on install

For more information, visit https://snapcraft.io/docs

<span style="color: gray; font-weight: bold">Examples:</span>
Install the latest snap from the stable channel, auto-connecting all applicable interfaces
  snap install btop 

Install a snap with classic mode, use the newest revision from the beta channel, and add it to a quota group
  snap install --beta --classic --quota-group=docker docker

Install a snap in development mode from a local file, under the specified name
  snap install --devmode --name jupyter-lab ./jupyter-lab.0.1.snap

<span style="color: gray; font-weight: bold">Related commands:</span>
  snap find               Find snaps to install
  snap info               Get information about a snap and its available channels
  snap refresh --hold     Hold a snap at the currently installed revision.
</pre>

<pre>
$ snapcraft help build
snapcraft 1.7.0
Usage:
  snapcraft build [&lt;option&gt;...] [&lt;part-name&gt;...]

Summary:
  Build artifacts defined for a part. All parts will be built if no part names are specified.

Global options:
  --help       Show this help message and exit
  --verbose    Show debug information and be more verbose
  --quiet      Only show warnings and errors, not progress
  --verbosity  Set the verbosity level to 'quiet', 'brief', 'verbose', 'debug' or 'trace'
  --version    Show the application version and exit

Build options:
  --debug          Shell into the environment if the build fails
  --shell          Shell into the environment in lieu of the step to run
  --shell-after    Shell into the environment after the step has run

Target options:
  --platform       Set platform to build for
  --build-for      Set architecture to build for

Build runner options:
  --destructive    Build in the current host
  --use-lxd        Build in a LXD container
  --remote         Build remotely with launchpad service

Related commands:
  clean   Remove a part's assets
  pull    Download or retrieve artifacts defined for a part  
  stage   Stage built artifacts into a common staging area
  prime   Prime artifacts defined for a part
  pack    Create the final artifact
  try     Prepare a snap for "snap try"

For a summary of all commands, run 'snapcraft help --all'. 
For more information about a command, run 'snapcraft help <command>' or browse the snapcraft reference documentation.
</pre>

## Topical Help
<pre>
$ snapcraft help Account
snapcraft 1.7.0

Global options:
  --help       Show this help message and exit
  --quiet      Only show warnings and errors, not progress
  --verbose    More verbose output, repeat to increase detail
  --version    Show the application version and exit

Accounts in snapcraft refer to the store that is use to publish the snaps you are packing
using snapcraft. To access the Snap Store, go to https://snapcraft.io/store and either 
create a developer account or login with your developer account credentials.

For larger projects and ISVs, it might be useful to also create a brand account in the 
name of the project or company. 

Account commands:
  login               Log in to the Snap Store
  export-credentials  Log in to the Snap Store exporting credentials
  logout              Clear Snap Store credentials
  whoami              Get information about the current login

For a summary of all commands, run 'snapcraft help --all'. 
For more information about a command, run 'snapcraft help <command>' or browse the snapcraft reference documentation.

</pre>