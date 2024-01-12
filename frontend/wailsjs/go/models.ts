export namespace main {
	
	export class FinanceResult {
	    taxes: number;
	    owner: number;
	    profit: number;
	    opEx: number;
	
	    static createFrom(source: any = {}) {
	        return new FinanceResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.taxes = source["taxes"];
	        this.owner = source["owner"];
	        this.profit = source["profit"];
	        this.opEx = source["opEx"];
	    }
	}

}


export namespace main {
	
	export class FinanceResult {
	    taxes: number;
	    owner: number;
	    profit: number;
	    opEx: number;
	
	    static createFrom(source: any = {}) {
	        return new FinanceResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.taxes = source["taxes"];
	        this.owner = source["owner"];
	        this.profit = source["profit"];
	        this.opEx = source["opEx"];
	    }
	}

}

