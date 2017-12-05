# wordsnack-solver
word snack game solver

Finnish words: http://kaino.kotus.fi/sanat/nykysuomi/kotus-sanalista-v1.zip with https://creativecommons.org/licenses/by/3.0/deed.fi license.
```
curl -O http://kaino.kotus.fi/sanat/nykysuomi/kotus-sanalista-v1.zip
cat kotus-sanalista_v1/kotus-sanalista_v1.xml | awk -F"<st><s>" '{print $2}' | awk -F"<" '{print $1}' > words/finnish.txt
```