

import {CalculatePersonalFinance} from '../wailsjs/go/main/App';
import {main} from "../wailsjs/go/models";


export async function calculatePersonalFinance(income: number): Promise<main.FinanceResult>{
    return await CalculatePersonalFinance(income)
}