export namespace meta {
	
	export class Meta {
	    name: string;
	    description: string;
	    versionSum: number;
	    versionText: string;
	    createdAt: string;
	
	    static createFrom(source: any = {}) {
	        return new Meta(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.versionSum = source["versionSum"];
	        this.versionText = source["versionText"];
	        this.createdAt = source["createdAt"];
	    }
	}

}

