

## Installation on Ubuntu
```
git clone git@github.com:canonical/design-workshop.git exp01
cd exp01
git remote rename origit workshop
git remote add origin git@github.com:hartmutobendorf/testee.git
git branch -M main
git push -u origin main
```

```
workshop launch --verbose
```
Now copy the code shown, and login to github on your browser.

```
code --folder-uri vscode-remote://ssh-remote+workshop@10.4.75.58/projectP
```
