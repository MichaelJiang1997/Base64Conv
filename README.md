<h1>Base64 Converter</h1>
<b>轻松实现文件的Base64编码与解码</b>
<h2>1.Build from src</h2>
 <p>确认你的计算机已经安装Golang<br>
 控制台输入：<code>go version</code><br>
 应该有类似输出：
<code>go version go1.11.4 windows/amd64</code><br>
 控制台切换到源码文件所在目录执行
  <code>go build -o Base64Conv.exe .\main.go</code>
  </p>
  <h2>2.Download binary （windows）</h2>
  <p><a href="/releases">到rerlases下载！</a></P>
<h2>3.How to use?</h2>
<pre>
Base64 Converter v0.0.1<Michael Jiang:sencom.top>
file to Base64 code:
Base64Conv -mode=F2B -input=[filename] -output=[filename]
Base64 code to file:
Base64Conv -mode=B2F -input=[filename] -output=[filename]
</pre>
