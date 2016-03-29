$(function(){   
	$('.am-btn.am-btn-primary').click(function(){
		var name=$('#name').val();
		var ipAddress=$('#ipAddress').val();
		var username=$('#username').val();
		var password=$('#password').val();
		var path=$('#path').val();
		var port=$('#port').val();
		var remark=$('#remark').val();
			if(name==''){
				alert('工程名不能为空');
				name.focus();
			}else if(ipAddress==''){
				alert('IP地址不能为空');
			}else if(username==''){
				alert('用户名不能为空');
				username.focus();
			}else if(password==''){
				alert('密码不能为空');
				password.focus();
			}else if(path==''){
				alert('工程路径不能为空');
				path.focus();
			}else if(port==''){
				alert('端口号不能为空');
				port.focus();
			}else if(remark==''){
				alert('备注不能为空');
				remark.focus();
			}else{
				$.ajax({
					cache: true,
					type: "POST",
					url:"/project/save",
					data:$('#project').serialize(),
					async: false,
					error: function(request) {
						alert("Connection error");
					},
					success: function(data) {
						$("#commonLayout_appcreshi").parent().html(data);
					}
				});
			}
   })

});