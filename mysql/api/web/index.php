<?php
require_once __DIR__.'/vendor/autoload.php';

use Symfony\Component\HttpFoundation\Request;

$app = new Silex\Application();

$connection = mysql_connect('mysql', 'root', getenv('MYSQL_ENV_MYSQL_ROOT_PASSWORD'));

$app->get('/', function() {
    return "Welcome to paas-mysql API";
});

$app->post('/create', function(Request $request) use($app, $connection) {
    if (!$connection) {
        return $app->json(array(
            'message' => 'Could not connect to Mysql server.'
        ), 500);
    }

    $database = $request->get('database');

    if (!$database) {
        return $app->json(array(
            'message' => 'You should provide a database name.'
        ), 400);
    }

    $username = $database;
    $password = randomPwd(12);

    if (!($err = createDatabase($connection, $database))) {
        return $app->json(array(
            'message' => 'Could not create database "'.mysql_error($connection).'".'
        ), 500);
    }

    if (!($err = createUser($connection, $database, $username, $password))) {
        return $app->json(array(
            'message' => 'Could not create user "'.mysql_error($connection).'".'
        ), 500);
    }

    return $app->json(array(
        'database_name' => $database,
        'username'      => $username,
        'password'      => $password,
    ));
});

$app->run();

function randomPwd($length){
    $a = str_split("abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTUVWXY0123456789");
    shuffle($a);
    return substr(implode($a), 0, $length);
}

function createDatabase($connection, $database) {
    return mysql_query(
        "CREATE DATABASE IF NOT EXISTS `$database` DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;",
        $connection
    );
}

function createUser($connection, $database, $username, $password) {
    $err = mysql_query(
        // This will create user if it does not exists. http://stackoverflow.com/a/16592722/899205
        "GRANT ALL ON `$database`.* to '$username'@'%' identified by '$password';",
        $connection
    );

    if (!$err) {
        return $err;
    }

    return mysql_query(
        "SET PASSWORD FOR '$username'@'%' = PASSWORD('$password');",
        $connection
    );
}
