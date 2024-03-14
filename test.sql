/* 

   UPDATE ysb_ddhz dh
       SET dh.xgdjbh = v_applyno,
           NOTES     = '已转ERP批发申请单:' || v_applyno
     WHERE dh.djbh = p_applyno;
     
     */
     
     begin
proc_jzt_int_sale;
end;
     
     
                              '由药帮忙订单:' || ds.orderno || '自动产生',
     
     BEGIN PRO_YYC_MAKE_SYKH2(); proc_yyc2batsaleapp_auto(); END;
     
     
     
     
     
     
     
     CREATE OR REPLACE PROCEDURE proc_jzt_int_sale IS
  v_applyno t_batsaleapply_h.applyno%type;
begin
  FOR rec IN (select yjj_billid from JZT_ORDER_BASE_INFO where is_zx = '否') LOOP
    --生成批发申请单
    --定义单号
    v_applyno := 'YJJ' || SUBSTR('00000000' || rec.yjj_billid, -8);

--申请主表
INSERT INTO t_batsaleapply_h
  (applyno, --varchar2(40)
   billcode, --varchar2(10)
   compid, --number(10)
   vencusno, --number(10)
   vencusname, --5
   ownerid, --varchar2(10)
   reckonerid, --varchar2(10)  --部门编号
   
   subitemid,
   busno,  --9 moren 8888
   paytype,
   cashtype,
   
   saler, --yewu yuan 
   accchked, --财务审核状态，默认0
   invoicetype,
  -- addrid,
   whlgroupid,
   lastmodify,
   
   lasttime,
   status, --number  moren (0)
   -- createtime
   checkbit1,
   checkbit2,
   checkbit3,
   checkbit4,
   checkbit5,
   createuser,
   createtime,
   indentflag,
   account_date,
   PLATFORMNO,
   sum_whlprice,  --30
   NOTES
   ) --data
  SELECT v_applyno as applyno, --AS applyno,
         'ACB', -- AS billcode,
         b.compid,
         b.vencusno, --AS customid,
         b.vencusname, -- 5
         '01', --货主
         '01', --部门编号
         
         0, --子结算户
         8888, --业务机构
         '01', --付款账期
         '01', --付款方式
         830038, --业务员     
         0, --财务审核状态，默认0
         '1', --发票类型
         -- 10027, --运输地点
         1006, --价格组
         830038, -- zuizhong xiugai shijian 
         sysdate,
         0, --status 状态
         -- sysdate
         0,
         0,
         0,
         0,
         0,
         830038, --接口人员编码
         sysdate,
         0,
         sysdate,
         a.order_code,
         0.000000,   --30
        '由药九九订单:' || a.order_code || '自动产生'
    FROM JZT_ORDER_BASE_INFO a, t_vencus b
       WHERE ( a.COMPANY_NAME=b.VENCUSNAME and is_zx = '否' AND b.compid = 1 and a.yjj_billid = rec.yjj_billid )or  (a.THIRD_CUST_CODE = b.vencusno
     AND b.compid = 1 --对接企业编码
     and is_zx = '否' and a.yjj_billid = rec.yjj_billid);
    
  /* WHERE a.THIRD_CUST_CODE = b.vencusno
     AND b.compid = 1 --对接企业编码
     and is_zx = '否'
     and a.yjj_billid = rec.yjj_billid;
     */
     
   --  batsaleno
     
         --申请明细  APPLYNO
         
INSERT INTO t_batsaleapply_d
      (applyno,
       rowno,
       wareid,
       wareqty,
       checkqty,
       purprice,
       purtax,
       saleprice,
       whlprice,
       maxqty,
       midqty,
       avgpurprice,
       redeemsum,
       lastwhlprice)
       
      SELECT
       v_applyno AS applyno,
             rownum,
             b.wareid,
             a.ORDER_NUMBER,
             a.ORDER_NUMBER,
             b.lastpurprice,
             b.purtax,
             a.SETTLEMENT_PRICE,
             a.SETTLEMENT_PRICE,
             b.maxqty,
             b.midqty,
             b.lastpurprice,
             0,
             b.lastsaleprice
        from JZT_ORDER_DETAIL_INFO a, v_ware b, JZT_ORDER_BASE_INFO c
       where a.PROD_NO = b.warecode
         and a.ORDER_CODE = c.ORDER_CODE
         AND b.compid = 1 --对接企业编码
         and c.yjj_billid = rec.yjj_billid;
    commit;

/* 
    --自动审核申请单
    UPDATE t_batsaleapply_h
       SET status       = 1,
           execdate     = sysdate,
           checker1     = 168,
           checkbit1    = 1,
           sum_whlprice =
           (select sum(wareqty * whlprice)
              from t_batsaleapply_d
             where v_applyno = applyno)
     WHERE applyno = v_applyno;*/
     
     

    --回写状态
    update JZT_ORDER_BASE_INFO
       set is_zx = '是'
     where yjj_billid = rec.yjj_billid;
 commit;
 --commit;  --非常关键,order需要加递交事务

  END LOOP;
END;
