+++
date = "2008-01-23T12:09:00Z"
slug = "php_late_static_bindings"
title = "PHP: Late Static Bindings"
tags = ["php", "web design", "code"]
categories = ["blog", "programming"]
aliases = [
  "/blog/2008/01/23/php_late_static_bindings",
  "/post/php_late_static_bindings"
]
+++

New to PHP 5.3.0 will be **late static bindings**. Currently you can use the `self` keyword in PHP to access static methods of the current class, the new `static` keyword allows you to access static methods through the inheritance tree. This example from the PHP website explains best...

### Code

```php
<?php
class A {
    public static function who() {
        echo __CLASS__;
    }
    public static function test() {
        static::who(); // Here comes Late Static Bindings
    }
}
class B extends A {
    public static function who() {
         echo __CLASS__;
    }
}
A::test();
B::test();
```

### Output

> AB

There are more examples in the PHP manual about [Late Static Bindings](https://www.php.net/manual/en/language.oop5.late-static-bindings.php).
