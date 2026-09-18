import requests
from bs4 import BeautifulSoup

url = "https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements"
r = requests.get(url)
soup = BeautifulSoup(r.text, "html.parser")

el = (
    soup.find("aside", {"id": "main-sidebar"})
    .find("nav", {"class": "left-sidebar"})
    .find("div", {"class": "left-sidebar__content"})
    .find("ol")
)
cnt = 0
lst = None
for child in el.children:
    cnt += 1
    if cnt == 11:
        lst = child
        break
lst = lst.find("details").find("ol")
ans = []
for child in lst:
    a = child.find("a")
    ccnt = 0
    c1 = None
    for cchild in a:
        ccnt += 1
        if ccnt == 1:
            c1 = str(cchild.contents[0])[1:-1]
    if ccnt == 1:
        ans.append(c1)
with open("tag_list.txt", "w") as fout:
    fout.write("\n".join(ans))
