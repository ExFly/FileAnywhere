<img src="{{ .file_qr }}">
<br>
{{ .file_url }}
<br>
{{ .filename }}

<br>
<form action="/file/upload" method="post" enctype="multipart/form-data" >
　　　<input type="file" id="upload" name="upload" />
　　　<input type="submit" value="上传" />
</form>
{{ .msg_file_upload }}
