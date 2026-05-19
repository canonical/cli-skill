## Section 2: Verb Taxonomy and Aspect Classification

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|
| add | lifecycle | telic | yes | remove | add-cloud, add-model, add-unit, add-user |
| attach | transfer | telic | yes | detach | attach-storage, attach-resource |
| bind | mutation | telic | partial | — | bind |
| bootstrap | lifecycle | telic | no | — | bootstrap |
| cancel | execution | telic | no | — | cancel-task |
| change | mutation | telic | partial | — | change-user-password |
| create | lifecycle | telic | partial | remove | create-backup, create-storage-pool |
| debug | observation | atelic | no | — | debug-code, debug-hooks, debug-log |
| default | mutation | telic | partial | — | default-credential, default-region |
| destroy | lifecycle | telic | no | — | destroy-controller, destroy-model |
| detach | transfer | telic | yes | attach | detach-storage |
| diff | observation | telic | no | — | diff-bundle |
| disable | access | telic | yes | enable | disable-command, disable-user |
| download | transfer | telic | no | — | download-backup, download |
| enable | access | telic | yes | disable | enable-command, enable-user |
| export | transfer | telic | partial | — | export-bundle |
| find | observation | telic | no | — | find, find-offers |
| grant | access | telic | yes | revoke | grant, grant-cloud, grant-secret |
| help | observation | telic | no | — | help, help-action-commands, help-hook-commands |
| import | transfer | telic | partial | — | import-filesystem, import-ssh-key |
| info | observation | telic | no | — | info |
| integrate | transfer | telic | partial | remove-relation | integrate |
| kill | lifecycle | telic | no | — | kill-controller |
| login | access | telic | yes | logout | login |
| logout | access | telic | yes | login | logout |
| migrate | migration | telic | partial | — | migrate |
| move-to | mutation | telic | partial | — | move-to-space |
| offer | transfer | telic | yes | remove-offer | offer |
| refresh | mutation | telic | partial | — | refresh |
| register | access | telic | yes | unregister | register |
| reload | mutation | telic | no | — | reload-spaces |
| remove | lifecycle | telic | yes | add / add-like | remove-application, remove-cloud, remove-unit, etc. |
| rename | mutation | telic | partial | — | rename-space |
| resolved | execution | punctual | no | — | resolved |
| resume | mutation | telic | yes | suspend | resume-relation |
| retry | execution | punctual | no | — | retry-provisioning |
| revoke | access | telic | yes | grant | revoke, revoke-cloud, revoke-secret |
| run | execution | atelic/telic | no | — | run, exec |
| scale | mutation | telic | partial | — | scale-application |
| scp | transfer | telic | no | — | scp |
| set | mutation | telic | partial | — | set-constraints, set-credential, set-firewall-rule, set-model-constraints |
| show | observation | telic | no | — | show-action, show-application, show-cloud, etc. |
| ssh | execution | atelic | no | — | ssh |
| status | observation | atelic | no | — | status |
| suspend | mutation | telic | yes | resume | suspend-relation |
| switch | mutation | punctual | partial | — | switch |
| sync | execution | telic | no | — | sync-agent-binary |
| trust | access | telic | partial | — | trust |
| unexpose | mutation | telic | yes | expose | unexpose |
| unregister | access | telic | yes | register | unregister |
| update | mutation | telic | partial | — | update-cloud, update-credential, update-k8s, update-secret, update-storage-pool |
| upgrade | mutation | telic | partial | — | upgrade-controller, upgrade-model |
| whoami | observation | telic | no | — | whoami |

### Orphan noun commands (not verbs)

The following commands are noun-first or noun-only forms and do not fit the verb taxonomy above. They cover list, show, and configuration surfaces:

`actions`, `charm-resources`, `clouds`, `config`, `constraints`, `controller-config`, `controllers`, `credentials`, `dashboard`, `disabled-commands`, `documentation`, `firewall-rules`, `machines`, `model-config`, `model-constraints`, `model-defaults`, `model-secret-backend`, `models`, `offers`, `operations`, `regions`, `resources`, `secret-backends`, `secrets`, `spaces`, `ssh-keys`, `storage`, `storage-pools`, `subnets`, `users`, `version`

---

