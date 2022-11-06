import { atom } from "recoil";

export const moduleActive = atom({
  key: "moduleActive",
  default: 'templateList', // 'templateList','claimList','claimList','creatTempalte','offerClaims','setLink','revocation'
});

export const templateInfos = atom({
  key: "templateInfos",
  default: {}
});

export const activeDrawerState = atom({
  key: 'activeDrawerState',
  default: {
    
  }
})

export const routerNm = atom({
  key: "routerNm",
  default: '', // 'templateList','claimList','claimList','creatTempalte','offerClaims','setLink','revocation'
});