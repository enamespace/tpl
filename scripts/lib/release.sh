function release:generate_changelog() {

    git-chglog ${TAG_VERSION} > ./CHANGELOG/CHANGELOG-${TAG_VERSION}
    git add ./CHANGELOG/CHANGELOG-${TAG_VERSION#v}.md
    git commit -a -m "docs(changelog): add CHANGELOG-${TAG_VERSION#v}.md"
}