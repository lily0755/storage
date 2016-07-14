package controller

import (
    "github.com/gin-gonic/gin"
    "app/smoba/model"       
)

type GameHelper struct {

}

func (this *GameHelper) GetQipaInfo(c *gin.Context) {
    
    gameHelperModel := &model.GameHelperModel{}    

    date,err_date := gameHelperModel.GetQipaInfoDateRedisResult()
    if err_date != nil {
        c.JSON(200, gin.H{
            "Ret": 500,
            "Code": 500,
            "Msg": "get qipa date fail",
            "Data": []string{},
        })  
        return       
    }

    info,err_info := gameHelperModel.GetQipaInfoRedisResult(date)    
    if err_info != nil {
        c.JSON(200, gin.H{
            "Ret": 501,
            "Code": 501,
            "Msg": "get qipa info fail",
            "Data": []string{},
        }) 
        return       
    }

    c.JSON(200, gin.H{
        "Ret": 0,
        "Code": 0,
        "Msg": "success",
        "Data": info,
    })
}


func (this *GameHelper) GetFunnyInfo(c *gin.Context) {
    
    gameHelperModel := &model.GameHelperModel{}
    
    date,err_date := gameHelperModel.GetFunnyHeroDateRedisResult()
    if err_date != nil {
        c.JSON(200, gin.H{
            "Ret": 500,
            "Code": 500,
            "Msg": "get qipa date fail",
            "Data": []string{},
        })   
        return     
    }

    info,err_info := gameHelperModel.GetFunnyInfoRedisResult(date)    
    if err_info != nil {
        c.JSON(200, gin.H{
            "Ret": 501,
            "Code": 501,
            "Msg": "get qipa info fail",
            "Data": []string{},
        }) 
        return       
    }

    c.JSON(200, gin.H{
        "Ret": 0,
        "Code": 0,
        "Msg": "success",
        "Data": info,
    })
}





func (this *GameHelper) GetQipaInfoAll(c *gin.Context) {
    
    gameHelperModel := &model.GameHelperModel{}    

    date,err_date := gameHelperModel.GetQipaInfoDateRedisResult()
    if err_date != nil {
        c.JSON(200, gin.H{
            "Ret": 500,
            "Code": 500,
            "Msg": "get qipa date fail",
            "Data": []string{},
        })  
        return       
    }

    info,err_info := gameHelperModel.GetQipaInfoRedisResultAll(date)    
    if err_info != nil {
        c.JSON(200, gin.H{
            "Ret": 501,
            "Code": 501,
            "Msg": "get qipa info fail",
            "Data": []string{},
        }) 
        return       
    }

    c.JSON(200, gin.H{
        "Ret": 0,
        "Code": 0,
        "Msg": "success",
        "Data": info,
    })
}

