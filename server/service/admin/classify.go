package service

import (
	"errors"
	"github.com/imroc/req"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/service/blog"

	"github.com/gin-gonic/gin"
)

func EditClassify(r response.EditClassifyResponse) (err error, msg string, classify model.SysClassify) {
	db := global.GVA_DB
	blog.AddDynamic()
	classify.TypeName = r.TypeName
	classify.Icon = r.Icon
	classify.ParentID = r.ParentID
	//新增
	if r.ID == 0 {
		err = db.Where("type_name = ?", r.TypeName).Find(&classify).Error
		if err == nil {
			return err, "名称重复", classify
		}

		err = db.Create(&classify).Error
		if err != nil {
			return err, "新增失败", classify
		}
	} else {
		classify.ID = r.ID
		err = db.Where("ID=?", r.ID).Save(&classify).Error
		if err != nil {
			return err, "更新失败", classify
		}
	}
	return err, "success", classify
}
func DeleteClassify(c *gin.Context) (err error, msg string) {
	db := global.GVA_DB
	blog.AddDynamic()
	var classify model.SysClassify
	id := c.Query("id")
	err = db.Where("parent_id = ?", id).Find(&classify).Error
	if err == nil {
		return err, "存在子目录"
	}
	err = db.Where("ID = ?", id).Find(&classify).Error
	if err != nil {
		return err, "不存在该目录"
	}
	//先查是否有元素存在下级
	err = db.Where("ID = ?", id).Unscoped().Delete(&model.SysClassify{}).Error
	if err != nil {
		return err, "删除失败"
	}
	return err, "删除成功"
}

//获取阿里分类图标
func GetIconfontClassify() (msg string, data request.IconfontRequest, err error) {
	url := "https://www.iconfont.cn/api/project/detail.json?pid=1753589&t=1607856959819&ctoken=BxcsDXW7sQIf64nICtv86fEY"
	header := req.Header{
		"cookie": "cna=JLYCFahVO2MCAXOuZ1rwUYvf; EGG_SESS_ICONFONT=U8AXvqwdm-42-umGXGwgKq_Emj2wuVCkA87TjZ3dn6xm2T4whio3sIKoy4kjkuBSusLMQ-0MhcjWBE1FwhfGmLoJPlWKxBTqlWxtg8B7bbS-9O34Tv72IgPQ1r24oun0snRyyyZxO7i7toNSocei-5jvIs9s73f0VPfgkVMkmyH-feaqP3aPCVNA9AR2QPTUo2Nk2_NgXkV-ip8C9gAFXBu-p8FzrhCg559Rc8S18g9KW3SEf8kWRIqXZrDWX-Nb9YlgPqmKg7wWjFFLGdxZHnKLtqMZdORf--Dqh-KX_kXiK86MiyCtLFgSUKejTCZSj4wG2qfDijmyonSx4VLs1Rup1D-_asrbFOH1E3xH9qrO0-RLYf4B76ZiTjv9FRwMp_zgri_Nuxr3zRD3q42oImt_ArU6LuZmU0H_hrrdLDKoOjJ01SGdUJmMZD3Xlt7pRCqFVWZchyhuSP_YnZIVa_7mIl4Bge2y42sK3No5UDre4b4VwWXpw8TA-0B1t2KTkl0ltBOqQBh4zcrM5Kzpppp_QUUW6QBrxryGLjH5xz4G_OSK221RHGrRDEfWVUpTq9dmg7nkJiR-9d_abC6czByN9iMaxGhkZ_YfGABgvIJa0Qa3vXD5IjInrm79xPS-bugB5QoFNAdrBwQju69FngHutb5LgNpHaNpQkL5gpYcfOUyCQ5GOmu4C_k-YOPFhqpeN_cXsJH6XA96hl7RgTfu08-jNo_NSX4GrAwP5_iWY8ARtYl5knwXSDkuHVLJAU-kr0eZSadFqhH-sAyYPntmjcTgtwMVjS6oQ1E6xkO1nMeLkzQKqbURV_2L5dp2Y8JYO3AGDE5Wb9NGBUCUn4vbEcHtTI7o2yZXFq-73okEQJze4YBUcDJ_efsRbGjNAzEc4Dqe8CXZHdHOpt3JvttcIlH8LArWnajOb_U6FYqvZ8OAcFROUY0mZX6FMGF24q-AcdmdSMxOXp_UWToZAq-an4NRPO4zX3xWwI4vjZFSlkwL0EWFOqCba3rOcr1GQDKOFpR42kFt74c_Mas9IlR8qBfb1puzRC894NaqrdMEu5TOZ-7VNx6oJ7AzCfVsozOUUlVtmYOt76gXNqq6bnNAoBrxF3DPrdPouzNUBfbc=; trace=AQAAAETmVULl6wIAjvVBMZdqcntyzGAf; ctoken=BxcsDXW7sQIf64nICtv86fEY; u=5455681; u.sig=ZbSp54gm6uQY5WIPUAzkwbJtD-vFZhQXBb2qBczjaO0; xlly_s=1; isg=BIqKZmNZ4RDDUG4XzKZ1fswY23Asew7VOFeYnRTAAl1NxyiB_Ar45QFx1zMbN4Zt",
	}
	res, err := req.Get(url, header)
	res.ToJSON(&data)
	if data.Code != 200 {
		return "获取错误,请联系开源作者更新字体图标库配置", data, errors.New("错误")
	}
	return "获取成功", data, nil
}
