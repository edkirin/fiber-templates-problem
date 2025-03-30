<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
</head>

<body>
    <ul>
        <li><a href="/">home</a></li>
        <li><a href="/about">about</a></li>
    </ul>

    <div class="container">
        {{ block "content" . }}{{ end }}
    </div>

</body>

</html>