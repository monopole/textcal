# textcal

Start a plain text columnar calendar
file for planning multi-week events
like sprints or vacations.

Send the output to a file then edit the
file to add notes on the right.

Install:
```
go install github.com/monopole/textcal@latest
```

Usage:
```bash
textcal {monthCount} {lineSkipCount}
```
`monthCount` defaults to 3.

`lineSkipCount` defaults to 0.

Examples:

> ```bash
> $ textcal
>    2025  Su  Mo  Tu  We  Th  Fr  Sa
>     Mar  23  24  25  26 [27] 28  29
>     Apr  30  31   1   2   3   4   5
>           6   7   8   9  10  11  12
>          13  14  15  16  17  18  19
>          20  21  22  23  24  25  26
>     May  27  28  29  30   1   2   3
>           4   5   6   7   8   9  10
>          11  12  13  14  15  16  17
>          18  19  20  21  22  23  24
>          25  26  27  28  29  30  31
>     Jun   1   2   3   4   5   6   7
>           8   9  10  11  12  13  14
>          15  16  17  18  19  20  21
>          22  23  24  25  26  27  28
>     Jul  29  30   1   2   3   4   5
> ```

> ```bash
> $ textcal 4 1
>    2025  Su  Mo  Tu  We  Th  Fr  Sa
>     Mar  23  24  25  26 [27] 28  29
>
>     Apr  30  31   1   2   3   4   5
>
>           6   7   8   9  10  11  12
>
>          13  14  15  16  17  18  19
>
>          20  21  22  23  24  25  26
>
>     May  27  28  29  30   1   2   3
>
>           4   5   6   7   8   9  10
>
>          11  12  13  14  15  16  17
>
>          18  19  20  21  22  23  24
>
>          25  26  27  28  29  30  31
>
>     Jun   1   2   3   4   5   6   7
>
>           8   9  10  11  12  13  14
>
>          15  16  17  18  19  20  21
>
>          22  23  24  25  26  27  28
>
>     Jul  29  30   1   2   3   4   5
>
>           6   7   8   9  10  11  12
>
>          13  14  15  16  17  18  19
>
>          20  21  22  23  24  25  26
>
>     Aug  27  28  29  30  31   1   2
>
>           3   4   5   6   7   8   9
> ```
