
commitAll() {
    git add -A
    git ci -m "add author info"
    git push
}

commitAll


test() {
    go test -v github.com/ligenhw/goshare/blog -run TestBlogDetails
    go test -v github.com/ligenhw/goshare/orm -run TestReflect
    go test -v -count=1 github.com/ligenhw/goshare/orm -run TestOrm
}


"You have an error in your SQL syntax; check the manual that corr...+138 more"