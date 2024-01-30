# https://github.com/charmbracelet/vhs
Output demo.gif

Require echo

Set Shell "bash"
Set FontSize 12
Set Width 600
Set Height 440
Set Padding 20

Type "tm" Sleep 500ms Enter

Sleep 3s

Type "1:23:45" Sleep 500ms Enter

Sleep 3s

Type "1h23m45s" Sleep 500ms Enter

Sleep 3s

Type "1:23:45 + 1:30" Sleep 500ms Enter

Sleep 3s

Type "1:23:45 + 20s" Sleep 500ms Enter

Sleep 3s

Type "1:23 - :23 - ::1" Sleep 500ms Enter

Sleep 3s

Type "1:23:45 + 00:00:00.56" Sleep 500ms Enter

Sleep 3s

Type "0:30 * 3" Sleep 500ms Enter

Sleep 3s

Type "3m / 3" Sleep 500ms Enter

Sleep 3s

Type "(1m + 30s) / 3" Sleep 500ms Enter

Sleep 3s

Type "1h + (3m + 10s) * 3 - 0:05" Sleep 500ms Enter

Sleep 3s

Ctrl+D

Sleep 3s

Type "tm '0:40 + 1h - 5s'" Sleep 500ms Enter

Sleep 3s

Type "echo '0:30 - 2h' | tm" Sleep 500ms Enter

Sleep 5s
