
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