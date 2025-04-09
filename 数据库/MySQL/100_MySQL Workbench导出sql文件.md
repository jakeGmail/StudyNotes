步骤 1：打开MySQL Workbench
首先，您需要打开MySQL Workbench。如果您尚未安装MySQL Workbench，请前往MySQL官方网站下载并安装。

步骤 2：连接到数据库
在MySQL Workbench中，点击"Database"菜单，然后选择"Connect to Database"选项。在弹出的对话框中，填写数据库连接信息，包括主机名、端口号、用户名和密码。点击"Test Connection"按钮测试连接是否成功，然后点击"OK"按钮保存连接。

步骤 3：选择要导出的数据库
在连接成功后，您将在MySQL Workbench的左侧窗格中看到一个数据库列表。选择要导出的数据库，右键点击并选择"Set as Default Schema"选项，以将其设置为默认数据库。

步骤 4：选择导出选项
在MySQL Workbench的顶部菜单栏中，点击"Server"菜单，然后选择"Data Export"选项。在弹出的对话框中，选择"Export to Self-Contained File"选项，该选项将生成一个包含完整SQL语句的文件。

步骤 5：指定导出文件的名称和路径
在"Data Export"对话框中，选择要导出的文件的保存位置。您可以点击"Browse"按钮选择文件保存的目录和文件名。

步骤 6：导出SQL文件
点击"Start Export"按钮开始导出过程。MySQL Workbench将生成一个SQL文件，并将其保存到您指定的位置。


-----------------------------------
©著作权归作者所有：来自51CTO博客作者mob64ca12f15103的原创作品，请联系作者获取转载授权，否则将追究法律责任
mysqlworkbench导出sql文件
https://blog.51cto.com/u_16213435/7283832