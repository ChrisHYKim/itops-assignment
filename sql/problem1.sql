-- 요구사항:
-- 1. HOUSE B/L (HBL_NO) 를 기준으로 각 B/L 별 컨테이너 수량(CNTR 개수) 을 집계
-- 2. 가장 많은 수량을 가진 B/L 1건을 조회
-- 3. 컨테이너 수량은 CNTR_NO 기준으로 COUNT
-- 4. 동일 수량 시 ETD 빠른 순으로 우선 선택
-- 5. 결과 컬럼: HBL_NO, CNTR_COUNT, ETD
-- 6. 정렬: CNTR_COUNT DESC, ETD ASC

-- 여기에 SQL 쿼리를 작성하세요
-- 1. MST_HBL_NO 항목을 선택 후, 갯수만큼 테이블 출력
-- 2. CNTR 테아블과 JOIN 진행 후, 내림차순을 진행한 후, ETD 빠른순으로 진행한다.
SELECT 
    MST_HBL_NO,
    COUNT(CNTR.CNTR_NO) AS CNTR_COUNT,
    MST.ETD
FROM 
    FMS_HBL_MST
JOIN FMS_HBL_CNTR CNTR ON MST.HBL_NO = CNTR.HBL_NO
GROUP BY
    MST.HBL_NO, MST.ETD 
ORDER BY
    COUNT(CNTR.CNTR_NO) DESC
    MST.ETD ASC
FETCH FIRST 1 ROW ONLY;
