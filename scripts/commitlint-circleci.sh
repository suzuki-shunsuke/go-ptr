if [ -n "$CIRCLE_PULL_REQUEST" ]; then
  npx commitlint --from master --to $CIRCLE_BRANCH
fi
