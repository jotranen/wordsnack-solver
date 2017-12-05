# wordsnack-solver
word snack game solver

```
cat kotus-sanalista_v1.xml | awk -F"<st><s>" '{print $2}' | awk -F"<" '{print $1}' > sanat.txt
```