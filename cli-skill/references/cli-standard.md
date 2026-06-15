# Canonical CLI standards

|||||
|--- |--- |--- |--- |
|Index|DE013|||
|Title|Canonical CLI Standards|||
|Type|Author(s)|Status|Created|
|Standard|Hartmut Obendorf, Amy Pattanasethanon|Approved|Feb 1, 2024|
||Reviewer(s)|Status|Date|
||Gustavo Niemeyer|Approved|2025-08-12|


This is a list of standards which originate from both discussions with Canonical’s senior tech leads and Design best practice. All points listed here must be strongly considered before contributing code at Canonical. Individually, everything written here may seem trivial or not consequential, but every single guideline is a crucial building block for creating a great user experience for our products.

Ideally, this document would be accompanied by a library providing the described functionality for all the languages we are using but at present we must rely on the diligence of humans as we have a python implementation focused on craft-tools, but no other standard libraries.

If you spot any of the problems detailed here in our existing command-line interfaces, don’t be afraid to raise issues fixing them—the CLI is what most people will first encounter – and remember – about our products. 


# The Language of the Command Line
Canonical Design strives for freedom, reliability, preciseness, and collaborativeness, and we aim to reflect these values in all our user experiences. Command line interfaces (CLIs) are one of the primary means of interacting with Canonical products, and their user experience will shape the perception of Canonical in users’ minds.

How you interact with the CLI directly informs your perception of the capabilities of the product you are using. Your user experience can be greatly improved by a grammar that is adequate to the scope of the product (use the verb paradigm for simple cases, carefully amend the set of verbs used to communicate additional complexity and power) as it draws on a logical, easily learned and remembered choice of words.

Conversely, if you mix different grammar styles, or pick an uneven vocabulary, this will make it difficult to deduce functionality (“what will this command do?” can be easily triggered by using vastly different verbs, or by mixing different objects without allowing the user to differentiate between them).


## Grammar + Vocabulary 
A command-line tool’s primary interface is its name; often, it will take a parameter and/or flags (that can have their own parameters). More complex functionality often uses a *command* as the first argument.


```
cp -R     directory_1 directory_2
   [flag] [parameter] [parameter]

docker exec      -it my_container  /bin/bash
       [command] [flag with param] [parameter] 
```


Single-purpose command-line tools might not require a command to be specified, they are the command. Some examples are built-in shell commands such as ls, echo, but also tools such as less , rsync or awk.


### The standard grammar for Canonical command-line interfaces  

At Canonical, most of the commands we build are complex – they manipulate more than one type of object, and often cover more than one aspect of managing the behaviour of the software they interact with. With this standard Canonical command grammar, we strive to create a precise, minimal interface.

<a name="rule-commands-are-verbs"></a>
**Commands are verbs**. Every command that acts on a primary object of a command (e.g. snaps for snap, virtual machines for multipass) *must* be a verb. 

Choosing the right verb is not trivial: it needs to imply or trigger recall of the object type it refers to. And when a command acts on different object types, it needs to help the user differentiate between these types as they are not explicitly stated in the command (e.g. install or refresh a snap vs. login to a store).

<a name="rule-commands-are-logically-grouped"></a>
**Commands are logically grouped.** Not all commands are equal. Some commands act on the same type of objects or act in a logically coherent domain. We differentiate between these domains by grouping commands (e.g. build lifecycle vs. store management for snapcraft).

Ideally, verbs in one command group are implicitly connected with the object they act upon, semantically close to each other, and different from verbs used in other command groups.

<a name="rule-verb-noun-form"></a>
**When verbs alone are not sufficient **to distinguish between objects, use the **verb-noun** form (e.g. set-quota for snaps).

<a name="rule-shorthand-for-listing"></a>
**Use** foobars as a shorthand **for listing information** about all instances of a specific type of *secondary *object **instead of **list-foobar (e.g. snap services over snap list-services for listing services of a snap, but snap list for listing snaps).

<a name="rule-showing-state-shorthand"></a>
**For showing state** without changing it, use the shorthand status over show-status. For specific secondary objects, use foobar-status over show-foobar-status (e.g. snapcraft release-status), and foobar over show-foobar							.

<a name="rule-complex-commands-use-flags"></a>
Commands can become quite complex if a conceptual hierarchy is involved (e.g. a vm has networks that can be attached to a bridge). Consider using flags to prevent verb-noun-noun compounds, and e.g. use create-network --vmhost vmh1 nw1 over create-vmhost-network vmh1:nw1.

<a name="rule-choosing-right-verbs"></a>
Choosing the right verbs is hard. We consider it to be an essential part of your design work, and in order to get this right, we need to make sure that we find the verbs that best capture (and thus convey) what a command allows the user to do, and what effects it will have when completed. 


#### Commonly used commands

For the sake of consistency, here is a list of standard commands that are commonly used in our tools:


|||
|--- |--- |
|tool init|local initialization (host machine, within one snap)|
|tool bootstrap|distributed initialization (e.g. cluster, vm, cloud, second snap)|
|tool list|overview of all instances of primary type|
|tool foobars|overview of all instances of a secondary type|
|tool show <id>|details for one instance of the primary type (*)|
|tool foobar <id>|details for one instance of a secondary type|
|tool status|current tool state|
|tool foobar-status|current state of a specific object type|
|tool start/stop|services, instances, long-running processes (usually non-blocking)|
|tool enable/disable|primary type, feature|
|tool get/set/unset|getting, setting configuration and restoring-to-default|
|tool create-foo <id>|create a new instance of a secondary object, use flags to specify hierarchy|
|tool delete-foo <id>|delete an instance of a secondary object|
|tool update-foo <id>|use flags to specify state changes|
|tool help|show help for commands|
|tool version|show release version|



(*) you may prefer using info over show if the tool name already defines the object type, e.g. snap info foo (snap is both the type of the main object and the command name).

<a name="rule-all-commands-must-converge"></a>
**All** commands **must** converge on a grammar based on the above rules. 

<a name="rule-at-most-one-sublevel"></a>
When designing the grammar for a command that covers clearly distinct workflows and struggling to find a grammar that enables users to easily map actions and objects to a specific workflow, you can introduce a second command level. At most *one* sublevel may be used, and all commands in that second level must follow the rules outlined above. 

Alternatively, the tool  should be split into several different tools that have closely associated names.


```
# CMD gains a new dimension, managing clusters of foo instead of individual food
# use a (only one) secondary level
cmd cluster config --secret "$(cat secret.txt)" 
# ALTERNATIVELY create a new command that specifically acts on the new dimension
cmd-cluster config --secret "$(cat secret.txt)" 
```



### Every command matters

Every new command should be considered with all existing commands in mind. This will help to find the right wording, and the right logical group. 


### Conclusion: Keep it Simple (for your Users) 
Good grammar is **concise**, and **easy to learn** and **remember**. When deciding what grammar to use, you should consider a series of factors:



1. What are the **objects** you expose for interactions? This is perhaps the most important decision, so be sure to understand not only your own perspective but also that of your users.
2. How large is the **command set** you will deliver now and in the future? While adding a command is usually simple and smooth, changing the grammar after building a user base will be very difficult. You should choose a path that will hold against that future.
3. What are the **expectations** of your users, what is their existing mental model, and how might it need to change? Consider both internal and external reference points, and aim to align your decision with other products we are developing. 


## Parameters, Flags and Options 

Arguments can be added to a command line to **specify the object(s)** that an action will be performed on, and to **provide context to, or modify the action** to be performed. 

Command lines can take the following kinds of arguments:


### Positional Parameters 

Positional parameters are parameters that are identified only by their position following the specified action. You should **only** use positional parameters when there is** no doubt** about the meaning of each position; this can be because the action is clearly directional and maps in a natural way to the position of the parameters:


```
cp sourcefile destfile
```


Positional parameters are difficult for people to remember and use correctly, especially if they could be used interchangeably; do not use positional parameters unless the order is natural and easily memorizable:


```
ln <target> <the other target>
```


 \
Those rules still apply even if it is a single parameter: 


```
start apache2
docker pull ubuntu
snapcraft revisions <snap-name>
git remote show (implicit)
```



### Flags

Flags modify the performed action. They come in two varieties, short and long. Short flags start with a single dash (-) for a single character flag, and **must** **only** be used when the action being performed is both frequent *and* easy to imply from the context and the character used. Long flags start with a double dash (--), and are more descriptive and make it easier to imply their action.

<a name="rule-no-dual-flags"></a>
As a policy, **do not offer both short and long flags for the same action.** Typically this is done because the CLI was not carefully designed, and thus no consideration was given to whether the flag should be short or long.  Having both means users need to memorize both phrasings, even if they're used infrequently and might be hard to imply.

<a name="rule-single-char-flags-stackable"></a>
Single character flags **should** allow being stacked or combined.


```
apt upgrade -yadf
apt upgrade -y -a -d -f
```


<a name="rule-flags-not-order-dependent"></a>
Flags **must not** be dependent on ordering.


```
apt upgrade -a -d       # must have the same effect as
apt upgrade -d -a       # same arguments in different order
```



### Commonly Used Flags
<a name="rule-minimum-flags"></a>
All tools **should** at a minimum support these flags:


```
--help
```



### Flags with Values 

Flags often accept values to further specify the action:


```
snapcraft pack --verbosity debug
```


<a name="rule-flag-value-separation"></a>
Flag name and value **must** support separation by whitespace. \
Flag name and value **may** support separation by an equal sign. When there is an equal sign present, following whitespace is a separation marker for the next option/argument.


```
rsync -e "ssh -p 2222"
rsync --rsh="ssh -p 2222"
```



#### Flags accepting multiple arguments

To enable a flag to accept several values, there are two accepted patterns for the time being: it can be repeated, in which case the flag name must be singular; otherwise, it may accept a comma-separated list of values, in which case the flag name must be plural. For example:


```
tool do-foo --channel=edge --channel=beta
tool do-foo --channels=edge,beta 
```


Typically the first form is preferred when the value has a rich set of undefined characters, which would make choosing the separator hard or non-obvious, while the second form is better for cases where the vocabulary is strictly defined and composed of short words. **Do not provide both forms.**


#### Flags accepting key=value arguments

If you need to allow setting an open-ended set of parameters for a given command, it might be preferable to treat these as key=value arguments to a single flag, where `"="` is used to separate the key and value.


```
tool do-foo --bar key=value
```


Analogously, the flag should be singular if accepting single values, and plural if accepting multiple values. The key should behave exactly like the flag: use a singular for single values, and a plural for several values.


```
charmcraft release my-charm --channel=edge --resource foo=7 --resource bar=1

tool do-foo --bar a-key=1 --bar a-key=2 --bar a-key=3
tool do-foo --bar b-keys=1,2
# or
tool do-foo --bars a-key=1,a-key=2
tool do-foo --bars a-key=1,b-keys=2,3

juju deploy haproxy -n 2 --constraints spaces=dmz,^cms,^database
```



### Special case: - - as separator for command arguments 

If the command accepts a complete command line as an argument, "--" is to be used to separate the command and its flags from the positional command, as the command will stop parsing and allow the input to be interpreted verbatim:


```
multipass exec docker -- docker kill e1261a3214
multipass exec docker -- snapcraft --help
```


Here, the --help is passed to snapcraft rather than multipass.


# Feedback 

An important aspect of feedback is the help a command is showing when a user is seeking to understand what a command is capable of and how to phrase a valid command. We are developing guidance for formatting help, further details will be posted in .


## Errors, warnings and success messages 

Feedback is important to provide information about



* **Errors:** something has gone wrong that will lead to a non-successful completion, or abortion of the executed command
* **Warnings:** the command has detected a non-ideal state (of the environment or the processed data), or the unsuccessful execution of a non-vital step, or a suggested action to assure the continued success of the execution (eg a depreciation warning).
* **Success messages:** the command, or a step of the execution has completed successfully. 

<a name="rule-messages-human-readable"></a>
All messages should be human-readable (almost always using natural language or tabularized data), and as short and succinct as possible while considering the task at hand.


## Use of Color and Font Weight (bold)

The output of the CLI may use color to provide a clear visual hierarchy. That is, differentiating between more important and less important information. Never rely on color as the only mechanism conveying information. For example, do not rely on color as the single element to signal to the user that an action was successful - the message should be enough on its own. Also, be careful not to introduce colorful elements without care - it's easy to misuse the capability and introduce noise, and even hinder comprehension by calling attention to irrelevant information. 

<a name="rule-color-capability-detection"></a>
CLIs should only enable color when they can identify such capabilities in the output stream. When the capability isn't identified (e.g. the standard output is redirected to a filethe stdout or stderr stream) or when the NO_COLOR environment is set, you must not use color ([https://no-color.org/](https://no-color.org/)).


### Terminal Colors 

For CLIs, use only a limited set of colors to emphasize and differentiate information, even as many modern terminals support 256 or more colors – these are mainly useful when you want to use subtle shading, create the perception of depth or subtle visual hierarchy in a TUI (terminal user interface). 

ANSI colors are defined as


||||||||||
|--- |--- |--- |--- |--- |--- |--- |--- |--- |
|Last digit|0|1|2|3|4|5|6|7|
|color|black|red|green|yellow|blue|magenta|cyan|white|



Their exact rendering can vary with the color palette that the User defines for the terminal window; they might not render exactly the same as you see them, but they will render consistently for the User. Limit CLI output to ANSI colors, and use boldening as an additional way to provide visual hierarchy.

If the terminal supports Operating System Control (OSC) sequences (most do, e.g. gnome, xfce4, xterm, ghostty, iterm2), it is possible to query the background and foreground color from the Terminal. You should do this  to differentiate between “light mode” and “dark mode”, and optimize output for both cases: green on white and blue on black should be avoided for longer text or critical information. [https://gist.github.com/ThinGuy/76ca97fb7c035c62d8d444ce1fcb5617](https://gist.github.com/ThinGuy/76ca97fb7c035c62d8d444ce1fcb5617) 


## Tabular Data
Management of objects (machines, instances, packages, …) will often require processing of relational data. When rendering data to the output stream for users, tables are often used to structure the information.

<a name="rule-table-format"></a>
Format of tables

All tables should follow a standard format:



* The default width for a column delimiter is two spaces.
* Column headers are usually left-aligned.
* If present, column headers should use upper case (e.g. NAME, STATUS, etc) and bold font.
* **Do not** provide ASCII decorations, e.g. lines, to delimit columns.
* Show column headers by default but allow them to be hidden with --no-headers.
* If there is information that might be formatted strangely or not always be present, you may use a column called “NOTES” as a catch-all. Always have this column be the last one.
* Do not use spaces within table cells unless absolutely necessary. Spaces are reserved for delimiting columns, and their use within the cell makes it harder for processing (e.g. awk).
* For shorter information, prefer shorter column names (e.g. REV, not REVISION).

```
$ snap list
NAME              VERSION   REV   TRACKING       PUBLISHER         NOTES
amberol           0.10.3      30  latest/stable  alexmurray✪       -
android-studio    2023.1.1   148  latest/stable  snapcrafters✪     classic
arianna           23.08.3     37  latest/stable  kde✓              -
ascii-draw        0.2.0       66  latest/stable  nokse22           -
bare              1.0          5  latest/stable  canonical✓        base
beekeeper-studio  4.1.10     244  latest/stable  matthew-rathbone  -
blender           4.0.2     4300  latest/stable  blenderfoundati✓  classic
```




### Responsive output for tables 

It's hard to propose a good design for a CLI table without invoking good judgement. Tables should focus on key information that benefits from being seen together in a list, rather than individually in more detailed command output. Typically important "primary key" information will be on the leftmost columns, while more repetitive or less relevant details will be on the right side.

If possible, keep rows under 80 characters, but sometimes that's too limiting for the information that needs to be displayed. Going over that is acceptable if justified, but care should be taken as tables are not a good place for long information, because it breaks the columns that justify the data being seen together in the first place, and when wrapping it loses much of the benefit it being a table (it really isn't anymore, visually).

As a useful pattern, when it's tempting to add too much information to a table, try harder to find the important information that benefits from being displayed that way, and offer a detailed display command that will more extensively display all data for individual items.


### Empty states for tables 

When the data for a table has zero entries in it, the output should be clear about the fact this is an empty table rather than an error to obtain the information.  For example:


```
$ snap list
No snaps installed.
```


<a name="rule-empty-state-stderr"></a>
This message should go into stderr to make it clear that this isn't the standard output of the command, but the exit code should still be zero as the command succeeded in its goal (listing any installed snaps).

Avoid this:


```
$ snap list
NAME  VERSION  REV  TRACKING  PUBLISHER  NOTES
```


<a name="rule-empty-state-machine-readable"></a>
As an exception, when the output is required to be machine-readable, headers must be included (and the empty state message omitted).


```
$ snap list | cat
NAME  VERSION  REV  TRACKING  PUBLISHER  NOTES
```


If the output is in a machine-readable format, the output should simply output the zero value. The following cases demonstrate an output with two elements, and then their equivalent output with zero elements.


```
$ foo list --format=json // Two items case
{item: "one"}
{item: "two"}

$ foo list --format=yaml // Two items case
items:
  - one
  - two

$ foo list --format=json  // Zero case
$ foo list --format=yaml  // Zero case
items: []

```



## Logging Output


### Timestamps 

For communicating exact time, use the date and time format defined in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601).


```
2024-06-29T03:24:20Z
```



### Verbosity

Different use-cases call for different levels of information; for command-line tools this is reflected in the verbosity level of the output. This may range from just a return code (no output stream) to concise output (only the most relevant information) to verbose (all the information that may be needed). In addition, more information about the internal processes of a command can be surfaced using common debug levels such as info, debug and trace.


#### The verbosity levels


|||||
|--- |--- |--- |--- |
|Mode|ShorthandFlag|Analogy|Detail|
|Quiet|--quiet||No output or success messages. Only error messages conveying the fact that the requested operation could not be performed should be shown.  For example, the quiet mode of grep may return code 0 or 1 depending on whether the string was found or not, but if there is an error in the expression, it will still show an error message pointing out the mistake.|
|Brief|--brief (default for CLI tools)|printf, syslog.notice|In brief mode, only critical information is informed to the user. This can include progress, and should include success or failure notifications. “Printf” is an analogy for what you show to a user/client in a way that is easily understandable or human-readable.|
|Verbose|--verbose (default for daemons)|logf, syslog.info|In verbose mode, you want to talk to the audiences in a more descriptive and thorough way about what happens as a command is executed. This can include information about transactions or a series of events. A user would want to see this mode when they want to see additional events or intermediate steps, either to gauge progress in more detail, or to understand where an execution is failing. “Logf” is used as an analogy to show that the mental model of when a user logs something, they expect to see a thorough explanation of the exchanges between client and machine. Messages in this mode should be written in proper, with appropriate capitalization and punctuation.|
|Debug|No shorthand flag|debugf, syslog.debug|For debug mode, think of two developers talking to one another trying to understand where things go wrong. They will exchange messages that contain more information than normal users would care to read. Often, specific knowledge about the internal execution steps is required to understand debug messages. Debug information is useful to understand why things are failing, or performing forensics on the produced logs in order to change internal configuration, or the logic of the tool itself.|
|Trace|No shorthand flag|trace|Tracing is a special mode that can provide additional information about where in the code structure the execution is at any time. Trace information is useful to analyze performance, measure how often code is executed, or understand how the code is being executed with live data. It provides you the information you would usually see in a debugger. The volume of data generated by this mode is going to be overwhelming for typical users, and it's often too much even for high level debugging purposes.|


### Ephemeral Feedback: To New Line or not to New Line

Output to the terminal can be erased and overwritten. That makes it possible to provide ephemeral feedback, i.e. feedback that will not be recorded on screen.

<a name="rule-ephemeral-tty-only"></a>
When the command outputs to a pipe or file, it should never try to rewrite its output, only when an interactive tty session is used.


|||
|--- |--- |
|Non-ephemeral feedback (on a new line)|Ephemeral feedback (overwriting the last line)|
|For some commands, the intermediate steps that are taken are of meaningful consequence and need to be known.  For example, a juju deploy kafka command could allocate a machine that the user is going to be paying for while it's alive. If that's an ephemeral message, it will be quickly overridden and the user might not notice that it took place.|For other commands, the intermediate steps taken are relevant to give some comfort to the user while they are waiting for the full operation to be performed, so they know where time is being spent. For example, in snap remove vlc, the final success outcome is well understood, and the intermediate steps are of less relevance.|



# CLI Copy and Tone of Voice 
CLIs and UIs must use consistent terminology. If a concept is e.g. a “machine” in a CLI, it must not be described using e.g. “node” in a UI, and vice versa. Every product should define where their source of truth lies, and align accordingly.


#### Be concise, precise, and clear at all times.

When “talking” to a user, a CLI command should strive to be 



* **Concise**: Stripped of excess. No clutter. 
* **Precise**: Use exact terminology. Provide critical details. 
* **Clear**: Unambiguous and immediately understandable.

```
$ sudo snap refresh
All snaps up to date.

$ snap refresh --hold=24h firefox
General refreshes of "firefox" held until 2025-07-26T14:10:53+01:00

$ snap list syf                                                                                                                                                                       error: no matching snaps installed
```



Avoid this:


```
# do not be chatty when you can be concise
$ foo refresh
We checked for the newest revisions for your installed snaps and you are all good!

# do not be vague when you can be precise
$ foo refresh --hold=24h firefox
Now holding refreshes for your snap.

# do not be ambiguous when you can be clear (here: showing no data) 
$ snap list syf                                                                                                                                                                       $
```



#### Use passive, succinct sentences when arguments or objects are missing.


```
Invalid URL: "htp://foo.bar"
```


Avoid this:


```
You need to provide a valid URL.
```



#### Use an active, direct, friendly tone of voice when explaining helpful information to the user, without being apologetic.


```
Create a new controller using "juju bootstrap" or connect to another controller that you have access to by using "juju register".
```


Avoid this:


```
Oops, something went wrong. Please either let juju create a new controller using "juju bootstrap" or have another controller connected that you have been given access to using "juju register".
```



#### Use “cannot” instead of “didn’t / couldn’t / wouldn’t / failed to / unable to / etc”.


```
error: cannot establish the connection
```


Avoid this:


```
error: connection couldn't be established
```



#### Do not use contractions.


```
error: cannot establish the connection
```


Avoid this:


```
error: can't establish the connection
```

