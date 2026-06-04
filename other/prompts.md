
## Prompt Export - 2026-05-19 (Session 1)

1. understand what the skill does and how it works, i want to extend it
2. run the skill on the juju repo
3. add instructions to the skill so that it installs toolchains on Ubuntu (preferring snap) if it is missing e.g. go, java, or uv
4. run the skill on the juju directory
5. can you export all my prompts in this session, and append them to /prompts.md please

## Prompt Export - 2026-05-19 (Session 2)

Context: Multi-agent evaluation of the cli-review skill. Ran the skill on qwen36-snap/ with 3 models (run2: Opus, GPT, Gemini), interviewed each about their experience, synthesized findings, then implemented the top fixes to SKILL.md, added a `give-feedback` command, restructured resources, and ran a second multi-agent evaluation (run3: 4 models including GPT-mini) to validate the changes.

1. what are the top 10 fixes?
2. implement 1-7
3. Add a command to the skill - give-feedback. This step should be run last, and include the interview instructions used above. For each of the high or medium suggestions, it should create an issue on github in the project canonical/cli-skill.
4. restructure the skill: * move deprecation and standard definition into a resources/ directory * move the definition of the feedback interview process into a new resources/feedback.md file
5. move standard/README.md to standard.md , and same for deprecation
6. i want you to run the full analysis ad command design steps on the qwen36-snap/ directory now. Do it in parallel with 4 sub-agents using different models - claude opus 4.6, gpt 5.4, and "gemini 3.1 pro (Preview)", "gpt 5 mini". Use a directory called like the agent to store the results for each sub-agent. do not reuse results of a previous run. After each agent has run the skill, interview them about their experience using the skill: Make sure their context is fresh. Start with asking them how they felt about the experience - what is the emotional state it left them in?, then go through the steps defined in resources/feedback.md. Write down a protocol of your interview with the agent in feedback.md, then analyse the answers and come up with insights in insights.md. Then report the top 10 issues to me, and create issues in github for canonical/cli-skill for each issue with a medium or high severity
7. can you export all my prompts in this session, and append them to /project/prompts.md please. make a new section, and provide some context about the session, concisely.

## Prompt Export - 2026-05-19 (Session 3)

Context: Ran the cli-review skill's full analysis and command design steps on juju/ with 2 models (GPT-5.4, Gemini 3.1 Pro), storing results in model-named directories under juju/.

1. i want you to run the full analysis ad command design steps on the juju/ directory now. Do it in parallel with 2 sub-agents using different models - gpt 5.4, and gemini 3.1 pro. Use a directory called like the agent to store the results for each sub-agenr
2. can you export all my prompts in this session, and append them to /project/prompts.md please

## Prompt Export - 2026-05-13 to 2026-05-19 (Session 4)

Context: Followed up on the "chunk the workflow" insight from agents/insights.md. Implemented phased workflow in SKILL.md (3 phases with checkpoints), updated run-agents.js system prompt, tested with GLM-5 on juju, then created comprehensive comparison.md scoring all 20 agent runs across both projects grouped by juju/qwen. Also created meta-comparison.md analyzing small-vs-large CLI performance differences. Moved JS helpers into scripts/.

1. Start implementation
2. test it
3. (terminal notification — agent run completed)
4. now check for all the results that are in agents/ juju/ and qwen36-snap/, and create a comparison of the performance of the different agents. create a comparison.md that details all kinds of interesting things that the agents have done similarly and differently. also calculate a score for each agent based of the findings percentage in each stage
5. add the analysis of all the directories in juju/runX and qwen36-snap/runX, and group the comparison by juju and qwen
6. move the js helpers into the scripts directory
7. now make a meta-comparison.md that shows differences in the performance for small (qwen) and large(juju) clis
8. can you export all my prompts in this session, and append them to /project/prompts.md please
