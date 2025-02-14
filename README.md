# Применение принципов SOLID

## Single Responsibility Principle

Каждый класс (структура) и модуль в проекте имеет одну четко определенную ответственность. Например, классы `Rabbit`, `Tiger`, и `Wolf` отвечают за представление конкретных животных и их поведения, а `Zoo` управляет коллекцией животных и инвентарем.

## Open/Closed Principle

Классы и модули открыты для расширения, но закрыты для модификации. Например, добавление нового типа животного не требует изменения существующих классов, а только добавление нового класса, который реализует соответствующие интерфейсы.

## Liskov Substitution Principle

Классы, реализующие интерфейсы, могут быть заменены их подтипами без изменения корректности программы. Например, `Tiger` и `Wolf` реализуют интерфейс `IPredator`, что позволяет использовать их взаимозаменяемо в коде, который работает с хищниками.

## Interface Segregation Principle

Интерфейсы разделены на более мелкие, специализированные интерфейсы, чтобы классы не были вынуждены реализовывать методы, которые они не используют. Например, интерфейсы `IAlive` и `IInventory` разделяют обязанности, связанные с животными и инвентарем.

## Dependency Inversion Principle

Модули верхнего уровня не зависят от модулей нижнего уровня, а оба зависят от абстракций. В проекте это достигается через использование интерфейсов и внедрение зависимостей, например, через конструкторы, как в случае с `Zoo` и `VetClinic`.
 