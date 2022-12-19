# Committed

Committed is a WYSIWYG Git commit editor that helps improve the quality of your
commits by showing you the layout in the same format as `git log`.

## Features

- Built-in editor
- Emoji support
- Switch to a different author profile
- Does not take over the entire screen
- Subject line counter
- Appends sign-off
- Formats body to a maximum of 72 characters
- Best practises recommendations

## Interface

### Main

The main view when the application is started.

```
commit 1234567890abcdef1234567890abcdef1234567890
Author: John Doe <john.doe@example.com>
Date:   Mon Jan 2 15:04:05 2006 -0700

    ┌───┐ ┌──────────────────────────────────────────────────────────────┐
    │ X │ │ Capitalized, short (50 chars or less) summary                │ 47/50
    └───┘ └──────────────────────────────────────────────────────────────┘

    ┌──────────────────────────────────────────────────────────────────────────┐
    │ More detailed explanatory text, if necessary.  Wrap it to about 72       │
    │ characters or so.  In some contexts, the first line is treated as the    │
    │ subject of an email and the rest of the text as the body.  The blank     │
    │ line separating the summary from the body is critical (unless you omit   │
    │ the body entirely); tools like rebase can get confused if you run the    │
    │ two together.                                                            │
    │                                                                          │
    │ Write your commit message in the imperative: "Fix bug" and not "Fixed    │
    │ bug" or "Fixes bug."  This convention matches up with commit messages    │
    │ generated by commands like git merge and git revert.                     │
    │                                                                          │
    │ Further paragraphs come after blank lines.                               │
    │                                                                          │
    │ - Bullet points are okay, too                                            │
    │                                                                          │
    │ - Typically a hyphen or asterisk is used for the bullet, followed by a   │
    │   single space, with blank lines in between, but conventions vary here   │
    │                                                                          │
    │ - Use a hanging indent                                                   │
    └──────────────────────────────────────────────────────────────────────────┘

      Signed-off-by: John Doe <john.doe@example.com>

Alt  + <enter> Commit <1> Author <2> Emoji <3> Summary <4> Body
Ctrl +     <c> Cancel
```

### Emoji

The emoji view reduces the position and lines of the body section to make space for a selector to appear. Filtering is available to narrow the choices down.

```
commit 1234567890abcdef1234567890abcdef1234567890
Author: John Doe <john.doe@example.com>
Date:   Mon Jan 2 15:04:05 2006 -0700

    ┌───┐ ┌──────────────────────────────────────────────────────────────┐
    │ X │ │ Capitalized, short (50 chars or less) summary                │ 47/50
    └─┬─┘ └──────────────────────────────────────────────────────────────┘
      └───────────────────────────────────┐
    ┌─────────────────────────────────────┴────────────────────────────────────┐
    │? Choose an emoji: █                                                      │
    │> x - Improve structure / format of the code.                             │
    │  x - Improve performance.                                                │
    │  x - Remove code or files.                                               │
    │  x - Fix a bug.                                                          │
    │  x - Critical hotfix.                                                    │
    │  x - Introduce new features.                                             │
    │  x - Add or update documentation.                                        │
    │  x - Deploy stuff.                                                       │
    │  x - Add or update the UI and style files.                               │
    └──────────────────────────────────────────────────────────────────────────┘

    ┌──────────────────────────────────────────────────────────────────────────┐
    │ More detailed explanatory text, if necessary.  Wrap it to about 72       │
    │ characters or so.  In some contexts, the first line is treated as the    │
    │ subject of an email and the rest of the text as the body.  The blank     │
    │ line separating the summary from the body is critical (unless you omit   │
    │ the body entirely); tools like rebase can get confused if you run the    │
    │ two together.                                                            │
    └──────────────────────────────────────────────────────────────────────────┘

      Signed-off-by: John Doe <john.doe@example.com>

Alt  + <enter> Commit <1> Author <2> Emoji <3> Summary <4> Body
Ctrl +     <c> Cancel
```

### Author

The author view moves the subject line down and reduces the height of the body
section. This provides space for a selector to choose the commit author.

```commit 1234567890abcdef1234567890abcdef1234567890
Author: John Doe <john.doe@example.com>
Date:   Mon Jan 2 15:04:05 2006 -0700

    ┌──────────────────────────────────────────────────────────────────────────┐
    │? Choose an author: █                                                     │
    │> John Doe <john.doe@example.com>                                         │
    │  John Doe <john.doe@example.org>                                         │
    └──────────────────────────────────────────────────────────────────────────┘

    ┌───┐ ┌──────────────────────────────────────────────────────────────┐
    │ X │ │ Capitalized, short (50 chars or less) summary                │ 47/50
    └───┘ └──────────────────────────────────────────────────────────────┘

    ┌──────────────────────────────────────────────────────────────────────────┐
    │ More detailed explanatory text, if necessary.  Wrap it to about 72       │
    │ characters or so.  In some contexts, the first line is treated as the    │
    │ subject of an email and the rest of the text as the body.  The blank     │
    │ line separating the summary from the body is critical (unless you omit   │
    │ the body entirely); tools like rebase can get confused if you run the    │
    │ two together.                                                            │
    │                                                                          │
    │ Write your commit message in the imperative: "Fix bug" and not "Fixed    │
    │ bug" or "Fixes bug."  This convention matches up with commit messages    │
    │ generated by commands like git merge and git revert.                     │
    │                                                                          │
    │ Further paragraphs come after blank lines.                               │
    │                                                                          │
    └──────────────────────────────────────────────────────────────────────────┘

      Signed-off-by: John Doe <john.doe@example.com>

Alt  + <enter> Commit <1> Author <2> Emoji <3> Summary <4> Body
Ctrl +     <c> Cancel
```

### Commit

Accepting the commit shows the output that will match the the `git log` command.

```
commit 1234567890abcdef1234567890abcdef1234567890
Author: John Doe <john.doe@example.com>
Date:   Mon Jan 2 15:04:05 2006 -0700

     X Capitalized, short (50 chars or less) summary

     More detailed explanatory text, if necessary.  Wrap it to about 72
     characters or so.  In some contexts, the first line is treated as the
     subject of an email and the rest of the text as the body.  The blank
     line separating the summary from the body is critical (unless you omit
     the body entirely); tools like rebase can get confused if you run the
     two together.

     Write your commit message in the imperative: "Fix bug" and not "Fixed
     bug" or "Fixes bug."  This convention matches up with commit messages
     generated by commands like git merge and git revert.

     Further paragraphs come after blank lines.

     - Bullet points are okay, too

     - Typically a hyphen or asterisk is used for the bullet, followed by a
       single space, with blank lines in between, but conventions vary here

     - Use a hanging indent

     Signed-off-by: John Doe <john.doe@example.com>

[master 1234567] Capitalized, short (50 chars or less) summary
 3 files changed, 2 insertions(+), 1 deletions(-)
```

## Installation

Install Committed with Homebrew.

```bash
  brew tap mikelorant/homebrew-custom
  brew install committed
```

## Shortcuts

**Global**

| Command            | Key       |
| :----------------- | :-------- |
| Commit             | alt-enter |
| Toggle sign-off    | alt-s     |
| Help               | alt-/     |
| Focus author       | alt-1     |
| Focus emoji        | alt-2     |
| Focus summary      | alt-3     |
| Focus body         | alt-4     |
| Cancel             | control-c |
| Next component     | tab       |
| Previous component | shift-tab |

**Emoji**

| Command       | Key       |
| ------------- | --------- |
| Clear emoji   | delete    |
| Reset filter  | escape    |
| Next page     | page down |
| Previous page | page up   |

## Architecture

The application uses the Bubbletea UI framework and is composed of nested
models. Simple boxes are combined with `textinput` and `textarea` models to
build the interface.

```
┌───────────────────────Model───────────────────────┐
│                                                   │
│ ┌─────────────────────Model─────────────────────┐ │
│ │ ┌───────────────────────────────────────────┐ │ │
│ │ │                                           │ │ │
│ │ │                   Info                    │ │ │
│ │ │                                           │ │ │
│ │ └───────────────────────────────────────────┘ │ │
│ └───────────────────────────────────────────────┘ │
│                                                   │
│ ┌─────────────────────Model─────────────────────┐ │
│ │ ┌───────────────────────────────────────────┐ │ │
│ │ │                                           │ │ │
│ │ │                  Header                   │ │ │
│ │ │                                           │ │ │
│ │ │ ┌───────────┐ ┌───────────┐ ┌───────────┐ │ │ │
│ │ │ │           │ │           │ │           │ │ │ │
│ │ │ │   Emoji   │ │  Summary  │ │  Counter  │ │ │ │
│ │ │ │           │ │           │ │           │ │ │ │
│ │ │ └───────────┘ └───────────┘ └───────────┘ │ │ │
│ │ └───────────────────────────────────────────┘ │ │
│ └───────────────────────────────────────────────┘ │
│                                                   │
│ ┌─────────────────────Model─────────────────────┐ │
│ │ ┌───────────────────────────────────────────┐ │ │
│ │ │                                           │ │ │
│ │ │                   Body                    │ │ │
│ │ │                                           │ │ │
│ │ └───────────────────────────────────────────┘ │ │
│ └───────────────────────────────────────────────┘ │
│                                                   │
│ ┌─────────────────────Model─────────────────────┐ │
│ │ ┌───────────────────────────────────────────┐ │ │
│ │ │                                           │ │ │
│ │ │                  Footer                   │ │ │
│ │ │                                           │ │ │
│ │ └───────────────────────────────────────────┘ │ │
│ └───────────────────────────────────────────────┘ │
│                                                   │
│ ┌─────────────────────Model─────────────────────┐ │
│ │ ┌───────────────────────────────────────────┐ │ │
│ │ │                                           │ │ │
│ │ │                  Status                   │ │ │
│ │ │                                           │ │ │
│ │ └───────────────────────────────────────────┘ │ │
│ └───────────────────────────────────────────────┘ │
│                                                   │
└───────────────────────────────────────────────────┘
```

## Git Commands

The following commands are used as reference to extract from Git specific fields
for the interface.

| Field                    | Method                                                       |
| :----------------------- | :----------------------------------------------------------- |
| commit hash              | indeterminable                                               |
| current branch           | `git rev-parse --abbrev-ref HEAD`                            |
| tip of branch            | `git log -1` = `git log HEAD -1`                             |
| upstream tracking branch | `git rev-parse --abbrev-ref ${branch}@{upstream}`            |
| user name                | `git config user.name`<br />`git config --global user.name`  |
| user email               | `git config user.email`<br />`git config --global user.email` |

## Authors

- [@mikelorant](https://www.github.com/mikelorant)

## License

[MIT](https://choosealicense.com/licenses/mit/)
