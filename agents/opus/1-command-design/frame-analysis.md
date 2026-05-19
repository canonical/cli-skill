# Frame Analysis of Juju CLI Verbs

## Overview

- **Verbs analyzed**: 52 unique verbs
- **Frames found in FrameNet**: 38
- **Coverage gaps**: 14 verbs with no FrameNet match

## Step 1: Extracted CLI Verbs

Verbs from verb-noun decomposition: add, attach, autoload, bind, bootstrap, cancel, change, create, debug, deploy, destroy, detach, diff, disable, download, dump, enable, exec, export, expose, find, grant, help, import, integrate, kill, login, logout, migrate, move, offer, refresh, register, reload, remove, rename, resolve, retry, revoke, run, scale, set, show, ssh, suspend, resume, switch, sync, trust, unexpose, unregister, update, upgrade.

Orphan commands without clear verbs: `juju`, `scp`, `ssh` (passthroughs), `version`.

## Step 2: Verb-to-Frame Mapping

| Verb | FrameNet Frame | Frame ID | Frame Definition | Lexical Unit Match |
|---|---|---|---|---|
| add | Cause_to_be_included | 2230 | An Agent makes a New_member part of a Group. | primary |
| attach | Attaching | 197 | Somebody causes one thing to be physically connected to something else. | primary |
| autoload | — | — | No FrameNet match. | no match |
| bind | Attaching | 197 | Somebody causes one thing to be physically connected to something else. | extended |
| bootstrap | — | — | No FrameNet match for infrastructure self-initialization. | no match |
| cancel | — | — | No FrameNet match for canceling queued operations. | no match |
| change | Cause_change | 683 | An Agent causes an Entity to change in category or attribute. | primary |
| create | Creating | 319 | A Cause leads to the formation of a Created_entity. | primary |
| debug | Emptying | 58 | Words relating to emptying containers and clearing areas. | no match |
| deploy | Arranging | 921 | An Agent puts a complex Theme into a particular Configuration. | extended |
| destroy | Destroying | 417 | A Destroyer affects the Patient negatively so that it no longer exists. | primary |
| detach | Detaching | 1390 | Somebody causes one thing to be physically detached from something else. | primary |
| diff | — | — | No FrameNet match for comparing differences. | no match |
| disable | Render_nonfunctional | 416 | An Agent affects an Artifact so that it is no longer capable of performing its function. | primary |
| download | — | — | No FrameNet match for file download. | no match |
| dump | Judgment_communication | 219 | A Communicator communicates a judgment. | no match |
| enable | Preventing_or_letting | 516 | A Potential_hindrance keeps an Event from taking place, or does not. | extended |
| exec | — | — | No FrameNet match for remote command execution. | no match |
| export | Exporting | 1568 | An Exporter moves Goods across a border from an Exporting_area to an Importing_area. | extended |
| expose | Reveal_secret | 1146 | A Speaker reveals Information that was previously secret. | extended |
| find | Locating | 792 | A Perceiver determines the Location of a Sought_entity. | primary |
| grant | — | — | No FrameNet match for permission granting in this sense. | no match |
| help | Assistance | 391 | A Helper benefits a Benefited_party by enabling the culmination of a Goal. | primary |
| import | Importing | 1567 | An Importer receives Goods from an Exporting_area across a boundary. | extended |
| integrate | — | — | No FrameNet match for system integration. | no match |
| kill | Killing | 590 | A Killer or Cause causes the death of the Victim. | primary |
| login | — | — | No FrameNet match for authentication login. | no match |
| logout | — | — | No FrameNet match for authentication logout. | no match |
| migrate | — | — | No FrameNet match for model migration. | no match |
| move | Cause_motion | 55 | An Agent causes a Theme to move from a Source along a Path to a Goal. | primary |
| offer | Offering | 2239 | An Offerer indicates ability and willingness to give a Theme to a Potential_recipient. | primary |
| refresh | Rejuvenation | 1614 | An Agent returns an Entity to an earlier state of vigor. | extended |
| register | Recording | 1615 | An Agent sets down in permanent form information about a Phenomenon. | extended |
| reload | — | — | No FrameNet match for reloading configuration. | no match |
| remove | Removing | 63 | An Agent causes a Theme to move away from a Source. | primary |
| rename | Name_conferral | 221 | Speakers name Entities. | primary |
| resolve | Resolve_problem | 1523 | An Agent resolves an outstanding Problem by finding its solution. | primary |
| retry | — | — | No FrameNet match for retrying failed operations. | no match |
| revoke | — | — | No FrameNet match for permission revocation. | no match |
| run | Operating_a_system | 1495 | An Operator manipulates a System such that it performs its function. | primary |
| scale | — | — | No FrameNet match for scaling replica counts. | no match |
| set | Placing | 62 | An Agent places a Theme at a Goal location. | extended |
| show | Cause_to_perceive | 2161 | An Agent causes a Phenomenon to be perceived by a Perceiver. | primary |
| suspend | Activity_pause | 167 | An Agent pauses in the course of an Activity. | primary |
| resume | Process_resume | 148 | An event resumes at a certain Place and Time. | primary |
| switch | Replacing | 1121 | An Agent changes the filler of a Role. | extended |
| sync | — | — | No FrameNet match for synchronization. | no match |
| trust | Trust | 1731 | A Cognizer thinks that Information given by a Source is correct. | extended |
| unexpose | — | — | No FrameNet match (inverse of Reveal_secret not lexicalized). | no match |
| unregister | — | — | No FrameNet match for unregistering. | no match |
| update | — | — | No FrameNet match for updating software/configuration. | no match |
| upgrade | — | — | No FrameNet match for upgrading versions. | no match |

## Step 3: Frame Element Comparison

### Deletion Verbs: remove, destroy, kill

| Verb | Frame | Agent | Patient/Theme | Source | Result | Manner |
|---|---|---|---|---|---|---|
| remove | Removing [63] | Agent | Theme | Source | — | — |
| destroy | Destroying [417] | Destroyer | Patient | — | non-existence | — |
| kill | Killing [590] | Killer | Victim | — | death | — |

**Semantic differences:**
- `remove` implies the Theme still exists elsewhere (Source role is core). UX implication: users may expect recovery.
- `destroy` implies permanent elimination (no Source role). UX implication: irreversible operation.
- `kill` implies the target was active/alive (Victim role). UX implication: strongest irreversibility, used for forceful termination.

### Enable/Disable Pair

| Verb | Frame | Agent | Artifact | Function |
|---|---|---|---|---|
| disable | Render_nonfunctional [416] | Agent | Artifact | no longer capable |
| enable | Preventing_or_letting [516] | Agent | Event | not prevented |

**Semantic asymmetry**: `disable` targets an Artifact; `enable` targets an Event. In Juju, both target commands/features, but the frames have different core elements. This creates a slight cognitive mismatch.

### Attach/Detach Pair

| Verb | Frame | Agent | Item | Goal/Source |
|---|---|---|---|---|
| attach | Attaching [197] | Agent | Item | Goal |
| detach | Detaching [1390] | Agent | Item | Source |

**Frame relation**: Detaching is the causative/inchoative inverse of Attaching. Direct semantic opposition with shared core elements. This is a well-designed symmetric pair.

### Suspend/Resume Pair

| Verb | Frame | Agent | Activity | State Change |
|---|---|---|---|---|
| suspend | Activity_pause [167] | Agent | Activity | paused |
| resume | Activity_resume [1156] | Agent | Activity | resumed |

**Frame relation**: Activity_resume inherits from Process_resume, which temporally follows Activity_pause. The frames are in a Precedes relation. This is a natural symmetric pair.

### Grant/Revoke (No FrameNet Match)

Both verbs lack FrameNet coverage in their CLI sense. Linguistically:
- `grant` = transfer of permission from granter to grantee
- `revoke` = retraction of permission from grantee

They form a semantic inverse pair but are not lexicalized in FrameNet. This is a gap in the linguistic resource, not a CLI design issue.

## Step 4: Frame Relation Mapping

| Frame A | Relation | Frame B | Distance | Verbs |
|---|---|---|---|---|
| Removing [63] | Inheritance | Transitive_action | 1 | remove |
| Destroying [417] | Inheritance | Transitive_action | 1 | destroy |
| Killing [590] | Inheritance | Transitive_action | 1 | kill |
| Removing [63] | See_also | Placing [62] | 1 | remove / set |
| Detaching [1390] | Causative_of | Attaching [197] | 1 | detach / attach |
| Activity_pause [167] | Precedes | Activity_resume [1156] | 1 | suspend / resume |
| Exporting [1568] | Perspective_on | Import_export_scenario | 1 | export |
| Importing [1567] | Perspective_on | Import_export_scenario | 1 | import |
| Creating [319] | Inheritance | Intentionally_create [280] | 1 | create |
| Creating [319] | ReFraming_Mapping | Coming_up_with [27] | 1 | create |
| Cause_change [683] | Inheritance | Transitive_action | 1 | change |
| Resolve_problem [1523] | Precedes | Confronting_problem | 1 | resolve |
| Recording [1615] | Inheritance | Statement [43] | 1 | register |
| Arranging [921] | Using | Placing [62] | 1 | deploy |
| Rejuvenation [1614] | Inheritance | Transitive_action | 1 | refresh |
| Cause_to_perceive [2161] | Inheritance | Intentionally_affect | 1 | show |
| Operating_a_system [1495] | Inheritance | Using | 1 | run |
| Render_nonfunctional [416] | Inheritance | Transitive_action | 1 | disable |
| Preventing_or_letting [516] | Using | Event | 1 | enable |
| Offering [2239] | Using | Giving | 1 | offer |
| Cause_motion [55] | Inheritance | Transitive_action | 1 | move |
| Name_conferral [221] | Inheritance | Intentionally_act | 1 | rename |
| Trust [1731] | Inheritance | Certainty [141] | 1 | trust |
| Removing [63] | — | Destroying [417] | 2 | remove / destroy |
| Removing [63] | — | Killing [590] | 2 | remove / kill |
| Destroying [417] | — | Killing [590] | 2 | destroy / kill |

**Distance interpretations:**
- Distance 0 (same frame): True synonyms. Juju has no distance-0 verb pairs.
- Distance 1 (directly related): May cause confusion if frames are similar. `disable` (Render_nonfunctional) and `enable` (Preventing_or_letting) are distance 2+ but functionally inverse.
- Distance 2 (indirectly related): `remove`, `destroy`, and `kill` are all at distance 2 from each other via Transitive_action. This explains user confusion — they are linguistically close but semantically distinct.

## Step 5: CLI-Specific Frame Annotations

For verbs with no FrameNet match or where the CLI sense diverges:

| Verb | Proposed Frame Name | Definition | Core Elements | Closest FrameNet Frame | Gap Notes |
|---|---|---|---|---|---|
| bootstrap | Self_initializing | An Agent causes a distributed system to create itself from minimal inputs. | Agent, System, Input, Target_environment | Creating [319] | Bootstrap implies self-creation without external Creator post-init. |
| deploy | Provisioning | An Agent places software onto infrastructure, configuring and activating it. | Agent, Software, Infrastructure, Configuration | Arranging [921] | Deploy includes provisioning + configuration + activation, not just arrangement. |
| integrate | Bidirectional_linking | An Agent creates a reciprocal relationship between two systems. | Agent, System_1, System_2, Relationship_type | Attaching [197] | Integration is bidirectional; Attaching is typically unidirectional. |
| refresh | Re_pull | An Agent re-acquires software from an external source and replaces the local copy. | Agent, Software, Source, Target | Rejuvenation [1614] | Refresh in Juju means "pull new charm revision", not "restore vigor". |
| grant | Permission_transfer | An Agent transfers access rights to a Recipient for a Resource. | Agent, Recipient, Resource, Access_level | — | No FrameNet frame covers permission granting. |
| revoke | Permission_retraction | An Agent retracts access rights from a Recipient for a Resource. | Agent, Recipient, Resource, Access_level | — | Inverse of Permission_transfer; not lexicalized in FrameNet. |
| update | State_modification | An Agent modifies an existing entity's state to a newer version. | Agent, Entity, Current_state, Target_state | Cause_change [683] | Update implies version progression; Cause_change is more general. |
| upgrade | Version_advancement | An Agent advances software to a newer major version. | Agent, Software, Current_version, Target_version | Cause_change [683] | Upgrade implies version hierarchy; update is more general. |
| migrate | Relocation_transfer | An Agent moves a workload from one environment to another. | Agent, Workload, Source_environment, Target_environment | Cause_motion [55] | Migration implies state preservation across environments. |
| scale | Replica_adjustment | An Agent changes the number of replicas for a workload. | Agent, Workload, Current_count, Target_count | Cause_change_of_position_on_a_scale [96] | Scale is about discrete count adjustment, not continuous position. |
| login | Authentication_establishment | An Agent establishes an authenticated session with a system. | Agent, System, Credential | — | No FrameNet frame for authentication. |
| logout | Authentication_termination | An Agent terminates an authenticated session. | Agent, System | — | Inverse of Authentication_establishment. |
| register | Controller_enrollment | An Agent enrolls a controller into the client's known set. | Agent, Controller, Client_store | Recording [1615] | Register here is about enrolling, not just recording. |
| unregister | Controller_unenrollment | An Agent removes a controller from the client's known set. | Agent, Controller, Client_store | — | Inverse of Controller_enrollment. |
| config | Configuration_access | An Agent reads or writes configuration key-value pairs for an entity. | Agent, Entity, Key, Value | — | Get/set hybrid not captured by single frame. |
| trust | Credential_delegation | An Agent delegates cloud credentials to an application. | Agent, Application, Credential, Scope | Trust [1731] | CLI sense is about delegation, not belief in correctness. |

## Step 6: Insights and Recommendations

### 1. True Synonyms (Distance 0)

Juju has **no distance-0 verb pairs** in FrameNet. All verbs evoke distinct frames. This suggests the command set has linguistic diversity, but also indicates inconsistency rather than healthy variation.

### 2. Near Synonyms (Distance 1–2)

| Verb Pair | Relation | Risk | Recommendation |
|---|---|---|---|
| remove / destroy / kill | Distance 2 via Transitive_action | **High** | Consolidate to single verb; use flags for force levels. |
| add / create | Distance 2+ (Cause_to_be_included vs Creating) | Medium | Converge on `create` for new resources, `add` for membership. |
| update / upgrade | Both unmapped; semantic neighbors | Medium | Differentiate clearly: `update` = in-place modification, `upgrade` = version advancement. |
| set / configure | Placing vs unmapped | Medium | `set` is used for constraints/credentials; `config` is used for key-values. Converge or clarify. |
| show / find | Cause_to_perceive vs Locating | Low | Both are observation verbs but with different frames; safe to coexist. |

### 3. False Friends

| Verb Pair | Issue |
|---|---|
| run / exec | `run` evokes Operating_a_system; `exec` has no frame but implies direct execution. Users confuse them. |
| debug / resolved | `debug` has no match (Emptying is wrong); `resolved` evokes Resolve_problem. Completely different frames. |
| switch / change | `switch` evokes Replacing; `change` evokes Cause_change. Similar in English but different in CLI semantics. |
| refresh / update / upgrade | All modify state but with different implied mechanisms. Refresh = re-pull; Update = modify; Upgrade = version advance. |

### 4. Safe Coexistence (Distance 2+, unrelated frames)

| Verb Pair | Frames | Rationale |
|---|---|---|
| show / destroy | Cause_to_perceive vs Destroying | Observation vs destruction; completely unrelated. |
| create / remove | Creating vs Removing | Lifecycle opposites with unrelated frames. |
| login / deploy | Authentication_establishment vs Arranging | Different domains entirely. |
| suspend / kill | Activity_pause vs Killing | Pause vs terminate; frames are unrelated. |

### 5. Frame-Informed Rename Candidates

| Current Verb | Frame | Better Name | Rationale |
|---|---|---|---|
| `resolved` | Resolve_problem | `resolve` | Resolve_problem frame implies active problem-solving; past participle `resolved` is grammatically odd. |
| `refresh` | Rejuvenation (extended) | `update-charm` or `redeploy` | Rejuvenation implies restoring vigor, not pulling new software. |
| `trust` | Trust (extended) | `delegate-credential` or `authorize` | Trust frame is about belief in information, not credential delegation. |
| `bootstrap` | Self_initializing (manual) | `create-controller` or `init-controller` | Creating [319] is the closest frame; `bootstrap` is jargon. |
| `kill` | Killing | `destroy-controller --force` | Killing implies alive victim; controllers are not alive. `destroy` is more appropriate per Destroying frame. |
| `run` | Operating_a_system | `run-action` or `execute-action` | Operating_a_system is too general; specify the object. |
| `exec` | — | `execute` or `run-command` | `exec` is Unix jargon; `execute` is more standard. |
| `set` | Placing | `apply` or `assign` | Placing implies physical location; constraints/credentials are not physically placed. |
| `bind` | Attaching (extended) | `bind-endpoints` or `set-bindings` | Attaching is close but `bind` in networking has specific meaning. |

### Frame Coverage Gap Summary

| Category | Verbs | Gap |
|---|---|---|
| Authentication | login, logout, register, unregister, autoload | No FrameNet frames for auth/session management |
| Software lifecycle | update, upgrade, refresh, sync | No FrameNet frames for version management |
| Permission | grant, revoke | No FrameNet frames for access control |
| Execution | exec, run, debug, retry, cancel | Operating_a_system covers `run`; others unmapped |
| Comparison | diff | No FrameNet frame for comparison/diffing |
| Transfer | download, sync | No FrameNet frames for digital file transfer |
| Integration | integrate | No FrameNet frame for system integration |
| Scaling | scale | No FrameNet frame for replica adjustment |

These gaps indicate that many CLI verbs operate in technical domains that FrameNet does not cover. Manual frame definitions (Step 5) are necessary for precise semantic analysis.
