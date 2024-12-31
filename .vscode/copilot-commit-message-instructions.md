Use these rules to generate a commit message in `Japanese` based on the user's input.

Generate a commit message that satisfies the following Conventional Commit configuration. Prompt the user for the following items in order, and use their input to generate the commit message:

1. type (required, select from the list below)
2. scope (optional)
3. subject (imperative, 64 characters or less)
4. body (optional)
5. breaking (indicate any breaking changes, if applicable)
6. issues (reference to related issue numbers, if applicable)
7. lerna (information for multiple package management, if applicable)

The commit message should follow this format:

{type}: {emoji}{subject}
(blank line)
{body}
(blank line)
BREAKING CHANGE: {breaking}
(blank line)
Closes: {issues}

The type must be one of the following, with the corresponding emoji added:

- test → 💍
- feat → 🎸
- fix → 🐛
- chore → 🤖
- docs → ✏️
- refactor → 💡
- style → 💄
- ci → 🎡
- perf → ⚡️
- release → 🏹

