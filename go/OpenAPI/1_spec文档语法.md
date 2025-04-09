

# 1 定义组件
```yml
components:
  schemas:
    CommercialAccount:
      type: object
      description: commercial account info in system
    Response:
      type: object
      description: response of api
      properties:
        code:
          type: integer
          format: int
          example: 1
          description: error code return by api
        data: {} # 定义data为任意类型，后续引用时，实例化为具体的类型
        msg:
          type: string
          example: OK
    OKResponse:
      type: object
      description: response ok of api
      properties:
        code:
          type: integer
          format: int
          example: 1
          description: error code return by api
        data:
          type: string
          example: OK
        msg:
          type: string
          example: OK
```


# 1 POST定义
## 1.1 定义body中的json结构
```yml
paths:
  /account/register:
  post:
    summary: Merchant Registration
    description: Register a new merchant.
    requestBody:
        required: true # 表示body必须存在
        content:
          application/json: # 定义body中的内容为json
            schema:
              properties:
                email: # 字段名
                  type: string # 字段类型
                  example: 77777@io.com # 参数的示例值
                phone: # 电话号码
                  type: string
                verify_code: # 验证码
                  type: integer
                  format: int # 进一步定义类型的具体格式
                password: # 密码
                  type: string
              required: # 定义json参数中必须存在的字段
                - verify_code
                - password
    responses:
      '200': # 定义返回200时的响应结构
        description: 成功响应信息描述
          content:
            application/json:
              schema:
                allOf:
                  - $ref: '#/components/schemas/OKResponse' # 引用其他组件
       '400':
          description: 失败响应信息描述
          content:
            application/json:
              schema:
                allOf:
                  - $ref: '#/components/schemas/Response'
                  - properties: # 描述'#/components/schemas/Response'组件中data的类型
                      data:
                        type: string
                        example: NOK
                    type: object
        
```

# 2 上传文件
```yml
  # 上传个人认证信息
  /account/authentication/personal:
    put:
      security:
        - BearerAuth: [ ]
      summary: Perform personal authentication.
      description: Upload personal authentication information.
      requestBody:
        content:
          multipart/form-data:
            schema:
              properties:
                commercial_id:
                  type: integer
                  format: int64
                  description: 实名的商户id
                real_name:
                  type: string
                  description: 真实姓名
                id_card_number:
                  type: string
                  description: 身份证号码
                id_card_front:
                  type: string
                  format: binary
                  description: 身份证正面图片
                id_card_back:
                  type: string
                  format: binary
                  description: 身份证反面图片
                id_card_by_hand:
                  type: string
                  format: binary
                  description: 手持身份证图片
              required:
                - real_name
                - id_card_number
                - id_card_front
                - id_card_back
                - id_card_by_hand
      responses:
        200:
          description: 上传实名信息成功
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Response"
```