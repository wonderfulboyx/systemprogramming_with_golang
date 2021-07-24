# メモ
## システムコール
- システムコールは「特権モードでOSの機能を呼ぶこと」
- 特権モードとは
    - CPUレベルで設定されている
    - 機能制限されないモード
- 特権モードでないアプリケーションからインターネット通信、ファイル入出力やメモリ割り当てを行うには？
    - OSが提供するシステムコールを使う。
    - システムコールが呼ばれると、その処理は特権モードで動作する。
- `os.File`構造体の内部ではsyscallパッケージの関数を呼び出し、システムコールをしている。
- goのランタイムは実行する必要のあるタスクに対して動けるスレッドが不足するとスレッドを立ち上げることがある。
- ファイルの読み書き、ネットワークアクセスはストレージのヘッドを動かす必要があり重い処理になる場合がある。メモリ確保もスワップが発生するとファイルの読み書きくらい遅くなる。

## POSIX
- Portable Operating System Interface
- IEEEで定められた規格
- C言語の関数名と引数、返り値が定義されている。
- golangのsyscallも先頭を小文字にすればC言語の関数名と同じ名前になっている。
- golangからはos.Fileなどを経由して触るのでふだん触る機会はない。ドキュメントもない。
- 触る必要がある場合は、C言語用の情報を参照する。

## syscallの中身
- `write` システムコールの実装 https://github.com/torvalds/linux/blob/v4.13/fs/read_write.c#L557-L572
- `SYSCALL_DEFINE3`マクロの展開部分 https://github.com/torvalds/linux/blob/9b76d71fa8be8c52dbc855ab516754f0c93e2980/include/linux/syscalls.h#L237-L254
- `asmlinkage long sys_write(...)` のように展開される。
- `asmlinkage`は引数をレジスタ経由で渡すようにするフラグ。スタックを使わずに情報が渡せるようになる。
- `sys_call` は `do_syscall_64` が呼び出している。
- `do_syscall_64` は `entry_SYSCALL_64` から呼ばれている。
- `entry_SYSCALL_64` を呼び出すのはCPU。

## golangとPOSIX
- POSIXはOS間のポータビリティを維持するためにC言語の関数を定義している。
- であれば、自前でsyscallを呼ぶのではなくC言語の関数をそのまま使うのが筋。
- しかし、golangは自前でポータビリティを頑張る選択をした。
- このアプローチにより、クロスコンパイルが容易になっている。

## モニタリング
- linuxでは`strace`を使う。
- macOSでは`dtruss`を使う。
    - ただし、OSのSIP(System Integrity Protection)を停止しないと動かない
    - goのコード側で`import "C"`しないと動かない
    

