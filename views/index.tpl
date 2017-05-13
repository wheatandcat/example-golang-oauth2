<!DOCTYPE html>

<html>

<head>
    <title>example-golang-oauth</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
    <div class="container">
        {{.Message}}
        <br />
        <br />
        <div class="well">
            oauth認証サンプル
        </div>
        <div class="panel panel-default">
            <div class="panel-heading">
                <i class="fa fa-google" aria-hidden="true" style="font-size:2.5rem;"></i>&nbsp;&nbsp;Googleログイン
            </div>
            <div class="panel-body">
                <a href="http://localhost:8080/google/oauth2">oauth2認証</a>
            </div>
        </div>
        <br />
        <div class="panel panel-default">
            <div class="panel-heading">
                <i class="fa fa-twitter" aria-hidden="true" style="font-size:2.5rem;"></i>&nbsp;&nbsp;twitterログイン
            </div>
            <div class="panel-body">
                <a href="http://localhost:8080/twitter/oauth">oauth認証</a>
            </div>
        </div>
        <br />
        <div class="panel panel-default">
            <div class="panel-heading">
                <i class="fa fa-facebook" aria-hidden="true" style="font-size:2.5rem;"></i>&nbsp;&nbsp;facebookログイン
            </div>
            <div class="panel-body">
                <a href="http://localhost:8080/facebook/oauth2">oauth2認証</a>
            </div>
        </div>
        <br />
        <div class="panel panel-default">
            <div class="panel-heading">
                <i class="fa fa-github" aria-hidden="true" style="font-size:2.5rem;"></i>&nbsp;&nbsp;githubログイン
            </div>
            <div class="panel-body">
                <a href="http://localhost:8080/github/oauth2">oauth2認証</a>
            </div>
        </div>
        <br />
    </div>
</body>

</html>
