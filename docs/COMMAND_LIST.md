# JUMPCHESS	COMMAND LIST

#### Success List

|  ID  |  command   | message |                             data                             |
| :--: | :--------: | :-----: | :----------------------------------------------------------: |
|  1   |   LOGIN    |  NONE   |            account : 用户账号    code : 用户密码             |
|  2   | LOGIN_DONE |  NONE   | id : 用户ID    username : 用户名    avatar : 头像    exp : 经验    level : 等级    coin : 金币    rank ： 段位    assert : 资产 |
|  3   |    BUY     |  NONE   |                     itemid : 购买商品ID                      |
|  4   |  BUY_DONE  |  NONE   |          itemid : 购买商品id    leftcoin : 剩余金币          |
|  5   |    PLAY    |  NONE   | select : 选择的角色id    potion : 选择的药水id    userid : 用户ID |
|  6   | PLAY_DONE  |  NONE   |                        room : 房间号                         |

#### Fail LIST

|  ID  |   command   |                   message                    | data |
| :--: | :---------: | :------------------------------------------: | :--: |
|  -1  | LOGIN_ERROR |    ACCOUNT_ERROR/CODE_ERROR/SERVER_ERROR     | NONE |
|  -2  |  BUY_ERROR  | COIN_NOT_ENOUGH/ALREADY_HAS_ONE/SERVER_ERROR | NONE |
|  -3  | PLAY_ERROR  |                 SERVER_ERROR                 | NONE |

