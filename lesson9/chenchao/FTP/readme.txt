
FTPserver
	���� go run FTPserver.go
	����-home  Ĭ��ΪbaseĿ¼�� �Զ���������Զ���rootĿ¼
	����server�ˣ����ʼ������homeĿ¼���ʹ�����־�ļ�

1��ʵ���� list   upload  get  
	list  ��ȡĿ¼�������ļ�
	upload  �ӿͻ��˽���һ���ļ�
	get	�ͻ�����������һ���ļ�

2���������־����
	test.log��¼�˿ͻ��˵�һЩ��������ʱ���¼



FTPclient

	go run FTPclient.go  -a  [list|upload|get] -n filename
	-a Ĭ��Ϊlist  -n Ĭ��Ϊall
	ִ�нű����Ӳ���Ĭ�ϻ�ȡĿ¼�������ļ�
1��list
	��ȡ��Ŀ¼�����е��ļ�   
2��upload
	��������ϴ�һ���ļ�
3��get
	����һ���ļ�
	