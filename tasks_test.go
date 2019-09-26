package aria2_test

import "testing"

func TestClient_AddUris(t *testing.T) {
	gid, err := client.AddUris("https://mirrors.tuna.tsinghua.edu.cn/anaconda/miniconda/Miniconda3-4.7.10-MacOSX-x86_64.pkg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("get gid ", gid)
	_, err = client.TellStatus("blablabla")
	if err == nil {
		t.Fatal("expecting error, but got ", err)
	}
	res, err := client.TellStatus(gid)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	_, err = client.Remove(gid, false)
	if err != nil {
		t.Fatal(err)
	}
}
