# Про Expirenza

## Один QR-код для всього

### Меню, оплата рахунку, чайові, бронювання та відгуки

Клікай на картинку

[![Видео](http://img.youtube.com/vi/etRjEB8gv50/0.jpg)](http://www.youtube.com/watch?v=etRjEB8gv50 "Видео")

## [Приклад реалізації базового функціоналу з використанням бібліотеки](../example/expirenza/simple)

### Принцип роботи

Сервіс Expirenza базується на принципі неперервного з'єднання ресторанного програмного комплексу з сервером Expirenza. 
Це з'єднання реалізовано через WebSocket, що забезпечує двосторонній обмін даними в реальному часі.  

Проте, унікальність цього сервісу полягає в тому, що обмін даними відбувається за принципом "WebSocket ініціює, ресторан відповідає через HTTP". 

Це означає, що всі запити від сервера до ресторану ініціюються через WebSocket, але відповіді від ресторану надсилаються через HTTP.  

Всі події, які можуть виникнути в процесі взаємодії, вже вбудовані в цю бібліотеку. 
Ваше завдання, як розробника, полягає в тому, щоб слухати ці події та реагувати на них відповідно. 
Відповіді, які ресторан повинен надсилати на кожну подію, вже визначені та прикріплені до відповідних подій.

#### Таблиця івентів та методи для відповідей

| Івент                            | Метод для відповіді         | Тіло/Структура відповіді              | Опис                                                                                                         |
|----------------------------------|-----------------------------|---------------------------------------|--------------------------------------------------------------------------------------------------------------|
| `GetBillEvent`                   | `GetBill`                   | `GetBillResponse`                     | очікує інформацію про рахунок.                                                                               |
| `PayBillRequest`                 | `PayBill`                   | `PayBillResponse`                     | очікує інформацію про оплату рахунку.                                                                        |
| `CreateOrderOnTableEvent`        | `CreateOrderOnTable`        | `CreateOrderOnTableResponse`          | створення та оплата замовлення на столі                                                                      |
| `CheckProductsRestrictionsEvent` | `CheckProductsRestrictions` | `CheckProductsRestrictionsResponse`   | очікує інформацію про обмеження продуктів                                                                    |
| `SplitOrderEvent`                | `SplitOrder`                | `SplitOrderResponse`                  | очікує інформацію про розділення рахунку                                                                     |
| `TablesInfoEvent`                | `TablesInfo`                | `TablesInfoResponse `                 | очікує інформацію про столи.                                                                                 |
| `UsersInfoEvent`                 | `UsersInfo`                 | `UsersInfoResponse`                   | очікує інформацію про персонал.                                                                              |
| `CategoriesInfoEvent`            | `CategoriesInfo`            | `CategoriesInfoResponse`              | очікує інформацію про категорії продуктів                                                                    |
| `HallPlansInfoEvent`             | `HallPlansInfo`             | `HallPlansInfoResponse`               | очікує інформацію про план залу.                                                                             |
| `CheckWorkEvent`                 | `CheckWork`                 | `CheckWorkResponse`                   | очікує інформацію про готовність до роботи.                                                                  |
| `SetSettingsEvent`               | `-`                         | `-`                                   | передача налаштувань, не очікує відповіді [подробиці](https://docs.expirenza.com/api/messages/setsettings)   |
| `StopListEvent`                  | `StopList`                  | `StopListResponse`                    | очікує інформацію про стоп-лист продуктів                                                                    |
| `VersionInfoEvent`               | `VersionInfo`               | `VersionInfoResponse`                 | очікує інформацію про версію плагіну                                                                         |
| `MenuInfoEvent`                  | `MenuInfo`                  | `MenuInfoResponse`                    | очікує інформацію про меню.                                                                                  |

