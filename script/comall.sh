
commitAll() {
    git add -A
    git ci -m "add author info"
    git push
}

commitAll


test() {
    go test -v github.com/ligenhw/goshare/blog -run TestBlogDetails
}