# CLI command set versioning and deprecation

**Abstract **This specification describes the process of deprecating or modifying CLI commands.

**Rationale **Humans and automated systems rely on stable command-line interfaces. You should avoid changing the command-line interface wherever possible as changes will disrupt the user's experience, and take a long time to roll out. Make sure you invest time into the initial design, so that it can deliver lasting value. When changes are necessary, provide a transition path, and minimize disruption to the user experience.

This document is describing how a change can be surfaced to Users of your command line. It should 


# Considerations for changing a command line interface


## Stability as a quality of command line interfaces (CLIs)

Command line tools are used in the terminal by humans. When a user intends to invoke a command, they are entirely dependent on their own memory or the externalized history of previous interactions with the command line interface. Changes in the command set can disrupt their workflow as the commands they correctly remember do not work anymore. They will need to learn what changes were made, and adapt their behavior.

Command line tools are also used in scripts. Once written, scripts rely on stable commands to successfully execute. When a command set is changed, a script will fail to have the intended effect. As subsequent commands may rely on the result of one successful execution, additional failures are likely. Worse, scripts may not provide feedback or relay failure when they fail to check whether a command was successfully executed, and simply assume success. 

This makes it highly desirable for a command line tool to provide a stable interface. And it raises the bar significantly for changes in command line interfaces to deliver value – minor improvements will often not deliver value as they are likely to cause more damage than benefit.

When you are considering to change a command line interface:



* Be aware that this is a very expensive change. The cost is not in changing the code, but in breaking users’ experience, and in having to support old conventions for deprecation purposes.
* Think of the people whose experience you are impacting. Will they be excited about the change, or will they be disappointed that you are breaking the workflow they have built around your tool?
* Small breaking changes are often not worth the effort, and a full refactoring is a major effort, and very difficult to get right. This will require a significant investment from your team, and your senior engineers.


## Reasons for changing a command line interface

While there are numerous reasons for stability, a command set will usually not stay unaltered forever. There is a need for change when



* new functionality is introduced
* existing functionality is removed
* the underlying conceptual and functional architecture had to be changed, and keeping the command bars user from understanding the new model 
* there are clearly identified problems with the current set of commands, or a clearly better alternative that is worth the cost of change

It is much easier to add functionality to a tool than to , and sometimes, this is even possible without a change to the command set.


## Types of changes to a command line interface


### Command set

CLIs can be changed on the command set level. There are two classes of changes, those that may not break the experience (adding a command) and those that are bound to break the experience, at least for some users (everything else).



* 
addition of a command


<p id="gdcalert1" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline drawings not supported directly from Docs. You may want to copy the inline drawing to a standalone drawing and export by reference. See <a href="https://github.com/evbacher/gd2md-html/wiki/Google-Drawings-by-reference">Google Drawings by reference</a> for details. The img URL below is a placeholder. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert2">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![drawing](https://docs.google.com/drawings/d/12345/export/png)

BREAKING CHANGE



* removal of a command
* renaming of a command
* changing the intended effects of a command
* reorganisation of the command hierarchy (if there are subcommands) 

The addition of a single command is unlikely to **break the interface** if the semantics of other commands are not affected. Introducing command after command still can **break the user experience** when the set of commands becomes too large to be easily remembered, or described; this might require you to plan ahead, or weigh the benefit and cost of introducing breaking structural changes.

**Removing, renaming **and** reorganizing** a command will **break the interface**. This should be avoided. If it needs to be done, proceed with intention and care, and follow the set of guidelines below.


### Flags

CLIs can also be changed on the flag level. There are two classes of changes, those that may not break the experience (adding a flag) and those that are bound to break the experience, at least for some users (everything else).



* addition of a flag



<p id="gdcalert2" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline drawings not supported directly from Docs. You may want to copy the inline drawing to a standalone drawing and export by reference. See <a href="https://github.com/evbacher/gd2md-html/wiki/Google-Drawings-by-reference">Google Drawings by reference</a> for details. The img URL below is a placeholder. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert3">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![drawing](https://docs.google.com/drawings/d/12345/export/png)

BREAKING CHANGE



* removal of a flag
* renaming of a flag
* change in the valid values or the grammar of a flag’s parameter
* changing the intended effects of a flag
* changing the order of arguments (if positional arguments are used in violation of [DE013](https://docs.google.com/document/u/0/d/1qsNbj1aipo6F6RuEqd3un43ArBoMenCiddCcrJBl23s/edit)) 

While the area of effect of changing a flag will likely be less serious than that of changing a command, the consequences are the same: adding a single flag will not break the interface, adding many flags can break the user experience. 

**Removing, renaming** and **changing** flag parameter expectations will **break the interface**. This should be avoided. If it needs to be done, proceed  with intention and care, and follow the set of guidelines below.


### Output

The output of a CLI can be changed, as well. It is less likely (but possible) that a change in the **help** of a command, in the **interactive reporting of progress**, or the **interactive prompting of input** is going to break compatibility.

Changing **machine-readable output** (e.g. JSON or YAML) is almost guaranteed to break compatibility, even if there are levels of impact: it is easier to process added data fields, whereas removing or renaming data fields is very likely to break applications that depend on their existence. Nonetheless, **these and all other changes** (e.g. error messages or return codes) in the output of a command are considered to** be breaking changes**. You can avoid breaking changes if you use versioning for machine-readable output.

Breaking changes should be avoided. If absolutely necessary, they should be  introduced only following the set of guidelines below.


# Guidelines for changing a command line interface

Changes in a command line interface need to follow clear rules that meet user expectations. User expectations may vary widely for different tools. We here use a “cycle” to define how long it takes to introduce new functionality, or communicate upcoming changes. For most of our tools, it will be close to an Ubuntu release cycle of 6 months, or the LTS cycle of 2 years, but depending on the environment, it can also be shorter or longer.


```
This document describes the minimal requirements to allow users to adapt to changes. If your project lives in an environment that places more value on stability, there are additional steps you should introduce to promote a change more gradually:
discuss upcoming changes publicly before committing to them
add an informational step, where you acknowledge the upcoming changes in the release notes (and debug or info level messages) first, announcing the version that will introduce the change
allow users to opt into the new interface, or opt out back to the previous interface for a time
go slower, e.g. breaking changes go into effect only after the earliest major version that is providing the old interface as a default goes out of support  
```


The considerations determining change for CLIs are similar to those for application programming interfaces (APIs), the rules  for change are therefore also similar to the [semantic versioning](https://semver.org/) used for most APIs. There are also similarities with how the application binary interface (ABI, the interface exposed by software that is defined for in-process machine code access) is evolved (e.g. for the [Ubuntu Kernel](https://wiki.ubuntu.com/KernelTeam/BuildSystem/ABI) or [GNU](https://gcc.gnu.org/onlinedocs/libstdc++/manual/abi.html)). 

Introducing changes should include at least the following steps, and respect the following timelines:


<table>
  <tr>
   <td><strong>Version</strong>
   </td>
   <td><strong>Type of Code Changes</strong>
   </td>
   <td><strong>Allowed Types of CLI Changes</strong>
   </td>
   <td><strong>CLI output change</strong>
   </td>
  </tr>
  <tr>
   <td><strong>Patch</strong> versions (x.x.<strong>N+1</strong>)
   </td>
   <td>fix bugs
   </td>
   <td>—
   </td>
   <td>fix output
   </td>
  </tr>
  <tr>
   <td><strong>Minor</strong> versions (x.<strong>N+1</strong>.x)
   </td>
   <td>improvements and refinements
   </td>
   <td>add command/flag
<p>
rename/move command/flag in help & docs but keep support for old command/flag
<p>
deprecation warning
   </td>
   <td>change only help and interactive output
   </td>
  </tr>
  <tr>
   <td><strong>Major </strong>versions (<strong>N+1</strong>.x.x)
<p>
>=1 cycle deprecation
   </td>
   <td>new functionality and breaking changes
   </td>
   <td>remove command/flag
<p>
following a rename/move, remove support for old command/flag
<p>
fail with message
   </td>
   <td>change non-machine readable output
<p>
rarely if never: change machine readable output
   </td>
  </tr>
  <tr>
   <td><strong>Next major</strong> versions (<strong>N+2</strong>.x.x)
   </td>
   <td>clean up legacy code
   </td>
   <td>remove failure message
   </td>
   <td>
   </td>
  </tr>
</table>


Allowed changes for commands and flags can also be represented in a timeline:



<p id="gdcalert3" ><span style="color: red; font-weight: bold">>>>>>  gd2md-html alert: inline image link here (to images/image1.png). Store image on your image server and adjust path/filename/extension if necessary. </span><br>(<a href="#">Back to top</a>)(<a href="#gdcalert4">Next alert</a>)<br><span style="color: red; font-weight: bold">>>>>> </span></p>


![alt_text](images/image1.png "image_tooltip")


*Figure 1: CLI command deprecation timeline*

**The interface must not be broken for minor releases. **The first release of a new command set needs to preserve compatibility with the previous versions. New commands and flags can be added, and renamed commands and flags can be changed in the help and documentation, but previous commands / flags must still work; the invocation of commands / flags that are now deprecated must emit a deprecation warning.

**A minimum of one cycle is required as a deprecation period **to allow users to adapt to changes. This means that the new interface needs to be stable for that period before the major version release. If your release environment is asking for more stability, you can also wait for a longer time to make sure users are aware of the changes.

**A failure message should be provided even after a new major version is released** as this provides information to script authors who have missed the deprecation notice in the minor versions leading up to the major version update. This message can be removed in subsequent major releases after at least one more cycle.

Messaging is an option only for changes in commands and flags. For **output changes**, there is **no way of communicating that change to the client**; it is therefore even more difficult to enable users to become aware of changes before they become breaking changes.


## Versioning and Store channels

A versioned upstream can be packaged in a store in different ways:



* using major version tracks, such as 5/stable, 6/stable, 7/stable
* using minor version tracks, such as 5.2/stable, 5.5/stable, 5.9/stable 
* (as an exception) using patch versions, such as 5.2.1/stable
* as a mix of the above.

Even if upstream should follow a different paradigm, we must provide stable interfaces: the stability guarantee applies to major tracks – **never break the CLI within a major version track**,  eg. 6/stable. It also covers minor version tracks if they exist, with the additional requirement that no change should become visible (and no command / flag should become deprecated).

Scripts are not to follow latest/stable or current/stable, but must always depend on a major version track.


## Versioning and Ubuntu series

Canonical supports specific Ubuntu LTS releases for 12+ years. The stability guarantee should apply to the LTS release. As a consequence, a major version shipped with an LTS release must be supported as long as the LTS series is supported. Further, all tracks released for a series must be supported as long as the LTS series is supported.

If you are planning to change the CLI of your tool, an ideal time would be following an LTS release as this gives you enough time to iterate and improve.


# Messaging: deprecating a command / flag


## Next minor version: deprecating a command / flag

When a command is to be removed from the command set, it should be deprecated in a minor version leading up to the release of a major version. This should ideally happen a full cycle before the release.

Deprecation means that the command / flag **still works**, and also that it emits a warning on stderr it will stop working with the next major release.


```
$ foo bar   # command is to be removed
warning: "foo bar" is deprecated and will be removed in the 3.0 release

$ foo bar   # command is to be renamed
warning: "foo bar" is deprecated, use "foo baz" instead
```


The same applies to removing a flag:


```
$ foo bar --qux   # flag is to be removed
warning: "--qux" flag is deprecated and will be removed in the 3.0 release

$ foo bar --qux   # flag is to be renamed
warning: "--qux" flag is deprecated, use "--waldo" instead
```


Documentation should reflect the deprecation: all references to a renamed command / flag should be changed to the new command / flag. All references to a command that is to be removed should be changed to show a notice that this command / flag is deprecated and will be removed in the next major release.


## Next major version: notification for a removed or renamed command / flag

When a command was removed, the execution should fail with the exit code 2. It should provide a clear error message on stderr. In the case of a recently removed or renamed command, this message should indicate that the command was removed or renamed.


```
$ foo bar   # command was removed
error: "foo bar" was removed in 3.0

$ foo bar   # command was renamed
error: "foo bar" was removed in 3.0, use "foo baz" instead
```


The same applies to removing a flag:


```
$ foo bar --qux   # flag was removed
error: "--qux" flag was removed in 3.0

$ foo bar --qux   # flag was renamed
error: "--qux" flag was removed in 3.0, use "--waldo" instead
```


Documentation should reflect the renaming / removal of a command in the changelog, and possibly also in an extra section “Removed / renamed commands”.


## Next+1 major version: clean up code and remove messaging

This notification should be presented during the lifetime of the next major release, or until the next LTS release; at this point, the notification can be removed and the code handling it can be cleaned up. Invocation of the command will now simply produce an error message on stderr.


```
$ foo bar   # command was removed or renamed
error: invalid command "foo bar"
```


The same applies to flags:


```
$ foo bar --qux   # flag was removed or renamed
error: invalid flag "--qux"
```


* \
Note: This should not prevent you from hinting at synonymous commands / flags if there is a user need.*


## An illustrative example

The LXD CLI tool lxc is being redesigned as part of the 25.10 and 26.04 cycles. As part of the redesign, flags that were using a mixture of long and short formats are being unified to only use one format.

This is a breaking change. It can only be rolled out in a major version, after one cycle of providing deprecation warnings.


```
$ lxc launch <image> <instance> -c key=value --config key=value       # Not OK
$ lxc launch <image> <instance> --config key=value --config key=value # OK
```


In this case, deprecation warnings are straightforward to define and implement:


```
# Minor version update introduces deprecation warning on stderr
$ lxc launch <image> <instance> -c key=value --config key=value
warning: "-c" flag is deprecated, use "--config" instead
[...]

# Next major version will fail with message on stderr
$ lxc launch <image> <instance> -c key=value --config key=value
error: "-c" flag was removed, use "--config" instead
[exit code 2]

# Subsequent major versions will fail with error on stderr
$ lxc launch <image> <instance> -c key=value --config key=value
error: invalid flag "-c"
[exit code 2]
```
