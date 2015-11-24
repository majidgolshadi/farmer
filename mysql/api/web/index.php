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

    if (!($err = createDB($connection, $database))) {
        return $app->json(array(
            'message' => "Could not create ".mysql_error()." database."
        ), 500);
    }

    if (!($err = defineUser($connection, $username, $password))) {
        return $app->json(array(
            'message' => "Could not user ".mysql_error()."."
        ), 500);
    }

    if (!($err = grantUserToDB($connection, $username, $database))) {
        return $app->json(array(
            'message' => "Could not set user privileges ".mysql_error()."."
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

function createDB($connection, $database_name) {
    return mysql_query(
        "CREATE DATABASE IF NOT EXISTS $database_name DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;",
        $connection
    );
}

function defineUser($connection, $username, $password) {
    mysql_query(
        "CREATE USER '".$username."'@'%' IDENTIFIED BY '".$password."';",
        $connection
    );

    if (1396 != mysql_errno()) {
        return "";
    }

    return mysql_query(
        "SET PASSWORD FOR '".$username."'@'%' = PASSWORD('".$password."');",
        $connection
    );
}

function grantUserToDB($connection, $username, $database) {
    return mysql_query(
        "GRANT ALL PRIVILEGES ON ".$database.".* TO '".$username."'@'%' WITH GRANT OPTION;",
        $connection
    );
}
